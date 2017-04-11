package logger

import (
	"path/filepath"
)

// GroupLogger print group period log
type GroupLogger struct {
	loggerMap     map[string]*PeriodLogger
	defaultLogger *PeriodLogger
}

// NewGroupLogger new a group logger manager
func NewGroupLogger(dir string, appName string, logSlice []string) *GroupLogger {
	gl := &GroupLogger{
		loggerMap:     map[string]*PeriodLogger{},
		defaultLogger: NewPeriodLogger(appName, "", dir, true),
	}

	for i := 0; i < len(logSlice); i++ {
		sliceName := logSlice[i]
		if len(sliceName) <= 0 {
			continue
		}
		downloadLogger := NewPeriodLogger(appName, sliceName, filepath.Join(dir, sliceName), false)
		gl.loggerMap["sliceName"] = downloadLogger
	}

	return gl
}

// L get the logger from group logger
func (gl *GroupLogger) L(group string) *Logger {
	d := gl.defaultLogger
	v, ok := gl.loggerMap[group]
	if ok {
		return NewLogger(d.Ls, v.Ls)
	}
	return NewLogger(d.Ls, nil)
}
