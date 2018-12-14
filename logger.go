package logger

import (
	"fmt"
	"log"
	"sync"
)

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
	depth int
}

// NewLogger new a logger
func NewLogger(defaultLogger *log.Logger, v *log.Logger, lv int) *Logger {
	return &Logger{
		d:     defaultLogger,
		v:     v,
		level: lv,
		depth: 2,
	}
}

// set depth
func (l *Logger) SetDepth(depth int) {
	l.depth = depth
}

// Print
func (l *Logger) Println(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&INFO <= 0 {
		return
	}
	l.d.SetPrefix("[INFO]")
	l.d.Output(l.depth, fmt.Sprintln(v...))
	if l.v != nil {
		l.v.SetPrefix("[INFO]")
		l.v.Output(l.depth, fmt.Sprintln(v...))
	}
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&INFO <= 0 {
		return
	}
	l.d.SetPrefix("[INFO]")
	l.d.Output(l.depth, fmt.Sprintf(format, v...))
	if l.v != nil {
		l.v.SetPrefix("[INFO]")
		l.v.Output(l.depth, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Print(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&INFO <= 0 {
		return
	}
	l.d.SetPrefix("[INFO]")
	l.d.Output(l.depth, fmt.Sprint(v...))
	if l.v != nil {
		l.v.SetPrefix("[INFO]")
		l.v.Output(l.depth, fmt.Sprint(v...))
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
	l.d.Output(l.depth, fmt.Sprintln(v...))
	if l.v != nil {
		l.v.SetPrefix("[DEBUG]")
		l.v.Output(l.depth, fmt.Sprintln(v...))
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&DEBUG <= 0 {
		return
	}
	l.d.SetPrefix("[DEBUG]")
	l.d.Output(l.depth, fmt.Sprintf(format, v...))
	if l.v != nil {
		l.v.SetPrefix("[DEBUG]")
		l.v.Output(l.depth, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&DEBUG <= 0 {
		return
	}
	l.d.SetPrefix("[DEBUG]")
	l.d.Output(l.depth, fmt.Sprint(v...))
	if l.v != nil {
		l.v.SetPrefix("[DEBUG]")
		l.v.Output(l.depth, fmt.Sprint(v...))
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
	l.d.Output(l.depth, fmt.Sprintln(v...))
	if l.v != nil {
		l.v.SetPrefix("[INFO]")
		l.v.Output(l.depth, fmt.Sprintln(v...))
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&INFO <= 0 {
		return
	}
	l.d.SetPrefix("[INFO]")
	l.d.Output(l.depth, fmt.Sprintf(format, v...))
	if l.v != nil {
		l.v.SetPrefix("[INFO]")
		l.v.Output(l.depth, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&INFO <= 0 {
		return
	}
	l.d.SetPrefix("[INFO]")
	l.d.Output(l.depth, fmt.Sprint(v...))
	if l.v != nil {
		l.v.SetPrefix("[INFO]")
		l.v.Output(l.depth, fmt.Sprint(v...))
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
	l.d.Output(l.depth, fmt.Sprintln(v...))
	if l.v != nil {
		l.v.SetPrefix("[WARN]")
		l.v.Output(l.depth, fmt.Sprintln(v...))
	}
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&WARN <= 0 {
		return
	}
	l.d.SetPrefix("[WARN]")
	l.d.Output(l.depth, fmt.Sprintf(format, v...))
	if l.v != nil {
		l.v.SetPrefix("[WARN]")
		l.v.Output(l.depth, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Warn(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&WARN <= 0 {
		return
	}
	l.d.SetPrefix("[WARN]")
	l.d.Output(l.depth, fmt.Sprint(v...))
	if l.v != nil {
		l.v.SetPrefix("[WARN]")
		l.v.Output(l.depth, fmt.Sprint(v...))
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
	l.d.Output(l.depth, fmt.Sprintln(v...))
	if l.v != nil {
		l.v.SetPrefix("[ERROR]")
		l.v.Output(l.depth, fmt.Sprintln(v...))
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&ERROR <= 0 {
		return
	}
	l.d.SetPrefix("[ERROR]")
	l.d.Output(l.depth, fmt.Sprintf(format, v...))
	if l.v != nil {
		l.v.SetPrefix("[ERROR]")
		l.v.Output(l.depth, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Error(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level&ERROR <= 0 {
		return
	}
	l.d.SetPrefix("[ERROR]")
	l.d.Output(l.depth, fmt.Sprint(v...))
	if l.v != nil {
		l.v.SetPrefix("[ERROR]")
		l.v.Output(l.depth, fmt.Sprint(v...))
	}
}

// fatal

// panic
