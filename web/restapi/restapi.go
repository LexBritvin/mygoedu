package restapi

import (
	"log"
	"mygoedu/myconfigurator"
	"net/http"
)

type DBlayerconfig struct {
	DB   string `json:"database"`
	Conn string `json:"connectionstring"`
}

func InitializeAPIHandlers() error {
	conf := new(DBlayerconfig)
	err := myconfigurator.GetConfiguration(myconfigurator.JSON, conf, "./web/apiconfig.json")
	if err != nil {
		log.Println("Error decoding config JSON", err)
		return err
	}
	h := NewCrewReqHandler()
	err = h.connect(conf.DB, conf.Conn)
	if err != nil {
		log.Println("Error connecting to DB", err)
		return err
	}
	http.HandleFunc("/api/crew/", h.handleCrewRequests)
	return nil
}

func RunAPI() error {
	if err := InitializeAPIHandlers(); err != nil {
		return err
	}
	return http.ListenAndServe(":8061", nil)
}
