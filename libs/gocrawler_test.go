package libs

import (
	"fmt"
	"github.com/lengsh/godingding/libs"
	"testing"
)

func TestGoCrawler(t *testing.T) {

	//libs.Test_Crawl()
	//libs.TestUpdate()
	//libs.CrawlMovieJob()
	//	fmt.Println(libs.CrawlStockJob("MSFT"))

	fmt.Println(libs.CrawlStockTestBaidu("BIDU")) //APPL"))
}
