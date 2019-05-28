package chat

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"sync"
	"testing"
	"time"
)

var once sync.Once

func chatServerFunc(t *testing.T) func() {
	return func() {
		t.Log("Starting chat server...")
		if err := Run(":2300"); err != nil {
			t.Error("Could not start chat server", err)
			return
		} else {
			t.Log("Started chat server")
		}
	}
}

func TestRun(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode...")
	}

	go once.Do(chatServerFunc(t))

	// Test clients after server starts.
	time.Sleep(1 * time.Second)

	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("Anonymous%d", rand.Intn(400))

	t.Logf("Hello %s, connecting to the chat system... \n", name)
	conn, err := net.Dial("tcp", "127.0.0.1:2300")
	if err != nil {
		t.Fatal("Could not connect to chat system", err)
	}
	t.Log("Connected to chat system")
	name += ":"
	defer conn.Close()
	msgCh := make(chan string)

	// Check sent message.
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			recvmsg := scanner.Text()
			sentmsg := <-msgCh
			if strings.Compare(recvmsg, sentmsg) != 0 {
				t.Errorf("chat message %s does not match %s", recvmsg, sentmsg)
			}
		}
	}()
	// Send messages
	for i := 0; i <= 10; i++ {
		msgbody := fmt.Sprintf("RandomMessage %d", rand.Intn(400))
		msg := name + msgbody
		_, err = fmt.Fprintf(conn, msg+"\n")
		if err != nil {
			t.Error(err)
			return
		}
		msgCh <- msg
	}
}

func TestServerConnection(t *testing.T) {
	t.Log("Test chat receive messages...")
	f := chatServerFunc(t)
	go once.Do(f)
	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:2300")
	if err != nil {
		t.Fatal("Could not connect to chat system", err)
	}
	t.Log("Connected to chat system")
	defer conn.Close()
}
