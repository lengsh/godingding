package libs

import (
	"fmt"
	"github.com/lengsh/godingding/libs"
	"testing"
)

func TestCrawler_Plugins(t *testing.T) {
	r := libs.Plugins{SoFile: "../so/stockplugin.so"}
	s := r.Crawler_Stock("baba")
	fmt.Println(s)

}
