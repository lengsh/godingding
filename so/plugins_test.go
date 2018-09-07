package libs

import (
	"fmt"
	"github.com/lengsh/godingding/libs"
	"testing"
)

func TestCrawler_Plugins(t *testing.T) {

	p := libs.Plugins{"./stockplugin.so"}
	s := p.Crawler_Stock("baba")
	fmt.Println("result(timedate):")
	fmt.Println(s)

}
