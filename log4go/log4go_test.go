package log4go 

import (
	"fmt"
	"github.com/lengsh/godingding/log4go"
	"testing"
	"os"
)

func TestCrawler_Plugins(t *testing.T) {
 fmt.Println("")
 gloger :=  log4go.NewF("log")
 gloger.Info("this is a test")
 gloger.Open()
 gloger.Debug("Next log should not be writen")
 gloger.Close()
 gloger.Debug("this log can't be writen")
}

func TestLogNew(t *testing.T) {
 gloger :=  log4go.New( os.Stdout)
 gloger.Info("this is a test")
 gloger.Open()
 gloger.Debug("Next log should not be writen")
 gloger.Close()
 gloger.Debug("this log can't be writen")

 gloger.Open()
 gloger.Debugf("%d, %s",12, "Next log should not be writen")
 gloger.Errorf("%d, %s",12, "Next log should not be writen")


}
