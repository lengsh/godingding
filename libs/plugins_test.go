package  libs 

import (
		"testing"
		"fmt"
		"github.com/lengsh/godingding/libs"
       )


func TestCrawler_Plugins(t *testing.T){
s :=  libs.Crawler_Stock("baba")
   fmt.Println(s)

}
