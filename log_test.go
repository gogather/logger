package logger

import (
	"log"
	"testing"
	"time"
)

func Test_Log(t *testing.T) {
	logSlice := []string{"download"}
	gl := NewGroupLogger("log", "dus", time.Hour*24, logSlice, log.Ldate|log.Ltime|log.Lshortfile)
	gl.L("download").Printf("hello world")
	log.Printf("hello this is default log")
}
