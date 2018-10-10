package libs

import (
	"fmt"
	"github.com/lengsh/godingding/libs"
	"testing"
)

func TestStock_webdriver(t *testing.T) {
	//	libs.CrawlStocksJob()
	s, _ := libs.LastStock("baba")
	fmt.Println(s)
}

/*
func TestCrawler_Phanmojs(t *testing.T) {
	s := libs.Crawler_163("baba")
	ss := libs.Crawler_Futu("baba", s)
	fmt.Println(ss)
}
*/
