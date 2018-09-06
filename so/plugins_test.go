package libs

import (
	"fmt"
	"github.com/lengsh/godingding/libs"
	"testing"
)

func TestCrawler_Plugins(t *testing.T) {
	s := libs.Crawler_Stock("baba")
	fmt.Println(s)

}
