package main

import (
	"flag"
	"mygoedu/chat"
	"mygoedu/mylogger"
	"mygoedu/web/portal"
	"strings"
)

func main() {
	logger := mylogger.GetInstance()
	logger.Println("Starting web service")
	operation := flag.String("o", "w", "Operation: w for wen \n c for chat")
	flag.Parse()
	switch strings.ToLower(*operation) {
	case "c":
		err := chat.Run(":2100")
		if err != nil {
			logger.Println("could not run chat", err)
		}
	case "w":
		err := portal.Run()
		if err != nil {
			logger.Println("could not run web portal", err)
		}
	}
}
