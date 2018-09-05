package libs

import (
	"fmt"
	"github.com/lengsh/godingding/libs"
	"testing"
)

func TestCrawler_Phanmojs(t *testing.T) {
	s := libs.Crawler_Phantomjs("baba")
	fmt.Println(s)

}
