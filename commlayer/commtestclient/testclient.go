package main

import (
	"flag"
	"log"
	"mygoedu/commlayer"
	"mygoedu/commlayer/myproto"
	"strings"
)

func main() {
	op := flag.String("type", "", "Server (s) or client (c) ?")
	address := flag.String("addr", ":8080", "address? host:port")
	flag.Parse()

	switch strings.ToUpper(*op) {
	case "S":
		runServer(*address)
	case "C":
		runClient(*address)
	}
}

func runServer(dest string) {
	c := commlayer.NewConnection(commlayer.Protobuf)
	recvChan, err := c.ListenAndDecode(dest)
	if err != nil {
		log.Fatal(err)
	}
	for msg := range recvChan {
		log.Println("Received: ", msg)
	}
}

func runClient(dest string) {
	c := commlayer.NewConnection(commlayer.Protobuf)
	ship := &myproto.Ship{
		Shipname:    "TestShip",
		CaptainName: "TestCaptain",
		Crew: []*myproto.Ship_CrewMember{
			{
				Id:           1,
				Name:         "Test1",
				SecClearance: 5,
				Position:     "Pilot",
			},
			{
				Id:           2,
				Name:         "Test2",
				SecClearance: 4,
				Position:     "Tech",
			},
			{
				Id:           3,
				Name:         "Test3",
				SecClearance: 3,
				Position:     "Enginneer",
			},
		},
	}

	if err := c.EncodeAndSend(ship, dest); err != nil {
		log.Println("Error occured while sending message", err)
	} else {
		log.Println("Send operation successful")
	}
}
