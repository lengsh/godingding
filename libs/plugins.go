package libs

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"os"
	"plugin"
)

type Plugins struct {
	SoFile string
}

func (r Plugins) Crawler_Stock(stk string) string {
	result := ""
	if _, err := os.Stat(r.SoFile); os.IsNotExist(err) {
		fmt.Println(r.SoFile + "  File does not exist")
		return result
	}

	p, err := plugin.Open(r.SoFile)
	if err != nil {
		fmt.Println(err)
		return result
	}
	crawlerstock, err := p.Lookup("CrawlerStock")
	if err != nil {
		logs.Error(err)
		return result
	}
	result = crawlerstock.(func(string) string)(stk)
	return result
}
