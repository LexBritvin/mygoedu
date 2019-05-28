package mylogger

import (
	"log"
	"os"
	"sync"
)

type mylogger struct {
	*log.Logger
	filename string
}

var logger *mylogger
var once sync.Once

func GetInstance() *mylogger {
	once.Do(func() {
		logger = createLogger("mygoedulogger.log")
	})
	return logger
}

func createLogger(fname string) *mylogger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	return &mylogger{
		filename: fname,
		Logger:   log.New(file, "Mylog ", log.Lshortfile),
	}
}
