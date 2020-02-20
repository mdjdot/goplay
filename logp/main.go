package main

import (
	"log"
	"os"
)

func main() {
	logFile, _ := os.OpenFile("kklog", os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	logger := log.New(logFile, "kklog: ", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("app start")
	logger.Panicln(os.ErrNotExist)
	logger.Println("app stopped\n")
	os.Exit(0)
}
