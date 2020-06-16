package logger

import (
	"log"
	"os"
)

var Log *log.Logger

func init() {
	var logpath = "./debug.log"
	var file, err = os.Create(logpath)

	if err != nil {
		panic(err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)
}
