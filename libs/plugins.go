package libs

import (
	"fmt"
	"os"
	"plugin"
)

type Plugins struct {
	SoFile string
}

func (r Plugins) Crawler_Stock(stk string) string {
	result := ""
	if _, err := os.Stat(r.SoFile); os.IsNotExist(err) {
		fmt.Println("File does not exist")
		return result
	}

	p, err := plugin.Open(r.SoFile)
	if err != nil {
		fmt.Println(err)
		return result
	}
	crawlerstock, err := p.Lookup("CrawlerStock")
	if err != nil {
		fmt.Println(err)
		return result
	}
	result = crawlerstock.(func(string) string)(stk)
	return result
}
