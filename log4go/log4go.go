package log4go

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

type G4Log struct {
	Logger       *log.Logger
	DebugEnabled *bool
}

var g4Logger *G4Log

func init() {

}

func New(out io.Writer) *G4Log {
	glog := new(G4Log)
	glog.Logger = log.New(out, "", log.LstdFlags)
	t := false
	glog.DebugEnabled = &t
	return glog
}

func SetDefaultLoger(r *G4Log) {
	if r == nil {
		panic("can't use nil")
	}
	g4Logger = r
}

func (r *G4Log) OpenDebug() {
	*r.DebugEnabled = true
}
func (r *G4Log) CloseDebug() {
	*r.DebugEnabled = false
}

func NewF(fp string) *G4Log {

	if fp == "" {
		wd := os.Getenv("GOPATH")
		if wd == "" {
			//panic("GOPATH is not setted in env.")
			file, _ := exec.LookPath(os.Args[0])
			path := filepath.Dir(file)
			wd = path
		}
		if wd == "" {
			panic("GOPATH is not setted in env or can not get exe path.")
		}
		fp = wd + "/log/"
	}

	year, month, day := time.Now().Date()
	filename := "log." + strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
	err := os.MkdirAll(fp, 0755)
	if err != nil {
		panic("logpath error : " + fp + "\n")
	}

	f, err := os.OpenFile(fp+"/"+filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("log file open error : " + fp + "/" + filename + "\n")
	}

	return New(f)
}

func getCaller() (string, int) {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	return file, line
}

func (l *G4Log) Debug(arg ...interface{}) {
	if !*l.DebugEnabled {
		return
	}
	fn, line := getCaller()
	s := fmt.Sprintf("[DEBUG](%s,%d):", fn, line)
	arg = append([]interface{}{s}, arg...)
	l.Logger.Println(arg...)
}

func (l *G4Log) Debugf(format string, arg ...interface{}) {
	if !*l.DebugEnabled {
		return
	}
	fn, line := getCaller()
	s := fmt.Sprintf("[DEBUG](%s,%d):", fn, line)
	l.Logger.Printf(s+format, arg...)
}

func (l *G4Log) Info(arg ...interface{}) {

	fn, line := getCaller()
	s := fmt.Sprintf("[INFO](%s,%d):", fn, line)
	arg = append([]interface{}{s}, arg...)
	l.Logger.Println(arg...)
}

func (l *G4Log) Infof(format string, arg ...interface{}) {

	fn, line := getCaller()
	s := fmt.Sprintf("[INFO](%s,%d):", fn, line)
	l.Logger.Printf(s+format, arg...)
}

func (l *G4Log) Error(arg ...interface{}) {
	fn, line := getCaller()
	s := fmt.Sprintf("[ERROR](%s,%d):", fn, line)
	arg = append([]interface{}{s}, arg...)
	l.Logger.Println(arg...)
}

func (l *G4Log) Errorf(format string, arg ...interface{}) {
	fn, line := getCaller()
	s := fmt.Sprintf("[ERROR](%s,%d):", fn, line)
	l.Logger.Printf(s+format, arg...)
}

func Error(arg ...interface{}) {
	if (g4Logger != nil) && (g4Logger.Logger != nil) {
		g4Logger.Error(arg...)
	}
}

func Debug(arg ...interface{}) {
	if (g4Logger != nil) && (g4Logger.Logger != nil) {
		g4Logger.Debug(arg...)
	}
}

func Info(arg ...interface{}) {
	if (g4Logger != nil) && (g4Logger.Logger != nil) {
		g4Logger.Info(arg...)
	}
}
