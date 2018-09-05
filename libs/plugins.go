package libs

import (
	//"errors"
	"fmt"
	"plugin"
)

func Crawler_Stock(stk string) string {
	result := ""
	p, err := plugin.Open("./so/stockplugin.so")
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
