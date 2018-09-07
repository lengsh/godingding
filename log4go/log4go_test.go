package log4go

import (
	"fmt"
	"github.com/lengsh/godingding/log4go"
	"os"
	"testing"
)

func TestCrawler_Plugins(t *testing.T) {
	fmt.Println("")
	gloger := log4go.NewF("log")
	gloger.Info("this is a test")
	gloger.OpenDebug()
	gloger.Debug("Next log should not be writen")
	gloger.CloseDebug()
	gloger.Debug("this log can't be writen")
}

func TestLogNew(t *testing.T) {
	gloger := log4go.New(os.Stdout)
	gloger.Info("this is a test")
	gloger.OpenDebug()
	gloger.Debug("Next log should not be writen")
	gloger.CloseDebug()
	gloger.Debug("this log can't be writen")

	gloger.OpenDebug()
	gloger.Debugf("%d, %s", 12, "Next log should not be writen")
	gloger.Errorf("%d, %s", 12, "Next log should not be writen")

}
