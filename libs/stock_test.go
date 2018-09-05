package  libs 

import (
		"testing"
		"fmt"
		"github.com/lengsh/godingding/libs"
       )


func TestCrawler_Phanmojs(t *testing.T){
s :=  libs.Crawler_Phantomjs("baba")
	   fmt.Println(s)

}
