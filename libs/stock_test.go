package libs

import (
	"fmt"
	"github.com/lengsh/godingding/libs"
	"testing"
)

func TestCrawler_Phanmojs(t *testing.T) {
	s := libs.Crawler_163("baba")
	ss := libs.Crawler_Futu("baba", s)
	fmt.Println(ss)
}
