package logger

import (
	"log"
	"testing"
	"time"
)

func Test_Log(t *testing.T) {
	logSlice := []string{"download"}
	gl := NewGroupLogger("log", "dus", time.Hour*24, logSlice, log.Ldate|log.Ltime|log.Lshortfile, ERROR|INFO)
	gl.L("download").Println("hello world")
	log.Printf("hello this is default log")
	gl.D().Print("default logger output")
}
