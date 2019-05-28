package portal

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"mygoedu/dblayer"
	"mygoedu/dblayer/passwordvault"
	"mygoedu/myconfigurator"
	"mygoedu/web/restapi"
	"net"
	"net/http"
	"sync"
)

var myWebTemplate *template.Template
var historylog = struct {
	logs []string
	*sync.RWMutex
}{RWMutex: new(sync.RWMutex)}

func Run() error {
	var err error

	conf := struct {
		Filespath string   `json:"filespath"`
		Templates []string `json:"templates"`
	}{}

	err = myconfigurator.GetConfiguration(myconfigurator.JSON, &conf, "./web/portalconfig.json")
	if err != nil {
		return err
	}
	myWebTemplate, err = template.ParseFiles(conf.Templates...)
	if err != nil {
		return err
	}

	restapi.InitializeAPIHandlers()
	log.Println(conf.Filespath)
	fs := http.FileServer(http.Dir(conf.Filespath))
	http.Handle("/", fs)
	http.HandleFunc("/crew/", crewHandler)
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/chat/", chatHandler)
	http.Handle("/chatRoom/", websocket.Handler(chatWS))
	go func() {
		err = http.ListenAndServeTLS(":8062", "cert.pem", "key.pem", nil)
		log.Println(err)
	}()
	err = http.ListenAndServe(":8061", nil)
	return err
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	nameStruct := struct{ Name string }{}
	r.ParseForm()
	if len(r.Form) == 0 {
		if cookie, err := r.Cookie("username"); err != nil {
			myWebTemplate.ExecuteTemplate(w, "login.html", nil)
			return
		} else {
			nameStruct.Name = cookie.Value
			myWebTemplate.ExecuteTemplate(w, "chat.html", nameStruct)
			return
		}
	}

	if r.Method == "POST" {
		user := r.Form["username"][0]
		pass := r.Form["password"][0]
		if !verifyPassword(user, pass) {
			myWebTemplate.ExecuteTemplate(w, "login.html", nil)
			return
		}
		nameStruct.Name = user
		if _, ok := r.Form["rememberme"]; ok {
			cookie := http.Cookie{Name: "username", Value: user}
			http.SetCookie(w, &cookie)
		}
	}
	myWebTemplate.ExecuteTemplate(w, "chat.html", nameStruct)
}

func verifyPassword(username, pass string) bool {
	db, err := passwordvault.ConnectPasswordVault()
	if err != nil {
		return false
	}
	defer db.Close()
	data, err := passwordvault.GetPasswordBytes(db, username)
	if err != nil {
		return false
	}
	hashedPass := md5.Sum([]byte(pass))
	return bytes.Equal(hashedPass[:], data)

}

func crewHandler(w http.ResponseWriter, r *http.Request) {
	db, err := dblayer.ConnectDatabase("mysql", "root:root@/mygoedu")
	if err != nil {
		return
	}
	all, err := db.AllMembers()
	if err != nil {
		return
	}
	err = myWebTemplate.ExecuteTemplate(w, "crew.html", all)
	if err != nil {
		log.Println(err)
	}
}

func chatWS(ws *websocket.Conn) {
	conn, err := net.Dial("tcp", "127.0.0.1:2100")
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	historylog.RLock()
	for _, log := range historylog.logs {
		err := websocket.Message.Send(ws, log)
		if err != nil {
			return
		}
	}
	historylog.RUnlock()

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			message := scanner.Text()
			err := websocket.Message.Send(ws, message)
			if err != nil {
				return
			}
		}
	}()

	for {
		var message string
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			return
		}
		// Send message to TCP server.
		_, err = conn.Write([]byte(message))
		if err == nil {
			historylog.Lock()
			if len(historylog.logs) > 20 {
				historylog.logs = historylog.logs[1:]
			}
			historylog.logs = append(historylog.logs, message)
			historylog.Unlock()
		}
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	about := struct {
		Msg string `json:"message"`
	}{}
	err := myconfigurator.GetConfiguration(myconfigurator.JSON, &about, "./web/about.json")
	if err != nil {
		return
	}
	err = myWebTemplate.ExecuteTemplate(w, "about.html", about)
	if err != nil {
		log.Println(err)
	}
}
