package main

import (
	"github.com/lengsh/godingding/libs"
)

func Crawler_Stock(st string) string {
	return libs.Crawler_Phantomjs(st)
}
