package logger

import "testing"

func Test_Log(t *testing.T) {
	logSlice := []string{"download", "upload", "convert"}
	gl := NewGroupLogger("log", "dus", logSlice)
	gl.L("download").Printf("hello world")
}
