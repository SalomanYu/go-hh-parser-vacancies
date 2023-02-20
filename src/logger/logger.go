package logger

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	Log *log.Logger
)

func init() {
	logpath := "info.log"
	flag.Parse()
	file, err := os.Create(logpath)
	if err != nil {
		panic(err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	fmt.Println("Logfile: " + logpath)
}