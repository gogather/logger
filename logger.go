package logger

import (
	"fmt"
	"log"
	"sync"
)

// GROUP option for enable group log
var GROUP = true

const (
	INFO  = 1 // 0001
	DEBUG = 2 // 0010
	WARN  = 4 // 0100
	ERROR = 8 // 1000
)

// Logger a logger
type Logger struct {
	level int
	d     *log.Logger
	v     *log.Logger
	mu    sync.Mutex
}

// NewLogger new a logger
func NewLogger(defaultLogger *log.Logger, v *log.Logger, lv int) *Logger {
	return &Logger{
		d:     defaultLogger,
		v:     v,
		level: lv,
	}
}

// Print
func (l *Logger) Println(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&INFO <= 0 {
		return
	}
	l.d.SetPrefix("[INFO]")
	l.d.Output(2, fmt.Sprintln(v...))
	if GROUP {
		l.v.SetPrefix("[INFO]")
		l.v.Output(2, fmt.Sprintln(v...))
	}
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&INFO <= 0 {
		return
	}
	l.d.SetPrefix("[INFO]")
	l.d.Output(2, fmt.Sprintf(format, v...))
	if GROUP && l.v != nil {
		l.v.SetPrefix("[INFO]")
		l.v.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Print(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&INFO <= 0 {
		return
	}
	l.d.SetPrefix("[INFO]")
	l.d.Output(2, fmt.Sprint(v...))
	if GROUP && l.v != nil {
		l.v.SetPrefix("[INFO]")
		l.v.Output(2, fmt.Sprint(v...))
	}
}

// debug

func (l *Logger) Debugln(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&DEBUG <= 0 {
		return
	}
	l.d.SetPrefix("[DEBUG]")
	l.d.Output(2, fmt.Sprintln(v...))
	if GROUP {
		l.v.SetPrefix("[DEBUG]")
		l.v.Output(2, fmt.Sprintln(v...))
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&DEBUG <= 0 {
		return
	}
	l.d.SetPrefix("[DEBUG]")
	l.d.Output(2, fmt.Sprintf(format, v...))
	if GROUP && l.v != nil {
		l.v.SetPrefix("[DEBUG]")
		l.v.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&DEBUG <= 0 {
		return
	}
	l.d.SetPrefix("[DEBUG]")
	l.d.Output(2, fmt.Sprint(v...))
	if GROUP && l.v != nil {
		l.v.SetPrefix("[DEBUG]")
		l.v.Output(2, fmt.Sprint(v...))
	}
}

// info

func (l *Logger) Infoln(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&INFO <= 0 {
		return
	}
	l.d.SetPrefix("[INFO]")
	l.d.Output(2, fmt.Sprintln(v...))
	if GROUP {
		l.v.SetPrefix("[INFO]")
		l.v.Output(2, fmt.Sprintln(v...))
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&INFO <= 0 {
		return
	}
	l.d.SetPrefix("[INFO]")
	l.d.Output(2, fmt.Sprintf(format, v...))
	if GROUP && l.v != nil {
		l.v.SetPrefix("[INFO]")
		l.v.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&INFO <= 0 {
		return
	}
	l.d.SetPrefix("[INFO]")
	l.d.Output(2, fmt.Sprint(v...))
	if GROUP && l.v != nil {
		l.v.SetPrefix("[INFO]")
		l.v.Output(2, fmt.Sprint(v...))
	}
}

// warn

func (l *Logger) Warnln(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&WARN <= 0 {
		return
	}
	l.d.SetPrefix("[WARN]")
	l.d.Output(2, fmt.Sprintln(v...))
	if GROUP {
		l.v.SetPrefix("[WARN]")
		l.v.Output(2, fmt.Sprintln(v...))
	}
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&WARN <= 0 {
		return
	}
	l.d.SetPrefix("[WARN]")
	l.d.Output(2, fmt.Sprintf(format, v...))
	if GROUP && l.v != nil {
		l.v.SetPrefix("[WARN]")
		l.v.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Warn(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&WARN <= 0 {
		return
	}
	l.d.SetPrefix("[WARN]")
	l.d.Output(2, fmt.Sprint(v...))
	if GROUP && l.v != nil {
		l.v.SetPrefix("[WARN]")
		l.v.Output(2, fmt.Sprint(v...))
	}
}

// error

func (l *Logger) Errorln(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&ERROR <= 0 {
		return
	}
	l.d.SetPrefix("[ERROR]")
	l.d.Output(2, fmt.Sprintln(v...))
	if GROUP {
		l.v.SetPrefix("[ERROR]")
		l.v.Output(2, fmt.Sprintln(v...))
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&ERROR <= 0 {
		return
	}
	l.d.SetPrefix("[ERROR]")
	l.d.Output(2, fmt.Sprintf(format, v...))
	if GROUP && l.v != nil {
		l.v.SetPrefix("[ERROR]")
		l.v.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Error(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&ERROR <= 0 {
		return
	}
	l.d.SetPrefix("[ERROR]")
	l.d.Output(2, fmt.Sprint(v...))
	if GROUP && l.v != nil {
		l.v.SetPrefix("[ERROR]")
		l.v.Output(2, fmt.Sprint(v...))
	}
}

// fatal

// panic
