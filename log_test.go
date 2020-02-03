package logger

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func Test_Log(t *testing.T) {
	logSlice := []string{"download"}
	gl := NewGroupLogger("log", "dus", time.Hour*24, logSlice, log.Ldate|log.Ltime|log.Lshortfile, ERROR|INFO)
	lt := gl.L("download")
	fmt.Println(lt)
	gl.L("download").Println("hello world")
	log.Printf("hello this is default log")
	gl.D().Print("default logger output")
}

func TestNewPeriodLogger(t *testing.T) {
	var lsl = map[string]*PeriodLogger{}
	pl := NewPeriodLogger("dus", "period", "log", false, log.Ldate|log.Ltime|log.Lshortfile, ERROR|INFO)
	//lg:=NewLogger(nil,pl.Ls,15)
	//var lpl = pl
	lsl["period"] = pl

	var tpl = lsl["period"]

	var lgt = NewLogger(tpl.Ls, tpl.Ls, 15)

	lgt.Println("sdsdsd")
	time.Sleep(time.Second * 2)
	pl.SwitchLog("log", "dus", "switch_1")
	lgt.Println("ererer")
}
