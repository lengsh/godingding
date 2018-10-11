package libs

import (
	"fmt"
	"github.com/lengsh/godingding/libs"
	"testing"
)

func TestStock_webdriver(t *testing.T) {
	//	libs.CrawlStocksJob()

	libs.CrawlStocksJob()

	s, er := libs.LastStock("baba")
	if er != nil {
		fmt.Println(er)
	} else {
		fmt.Println(s)
		dingtalker := libs.NewDingtalker()
		dingtalker.SendChatTextMessage(s.String())
	}

}

/*
func TestCrawler_Phanmojs(t *testing.T) {
	s := libs.Crawler_163("baba")
	ss := libs.Crawler_Futu("baba", s)
	fmt.Println(ss)
}
*/
