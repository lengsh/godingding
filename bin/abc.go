package main

import (
	"fmt"
	"github.com/lengsh/godingding/libs"
	"os"
	"strings"
	"time"
)

func main() {
	stk := "BABA"
	if len(os.Args) > 1 {
		stk = os.Args[1]
		stk = strings.ToUpper(stk)
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	s := libs.Crawler_Futu(stk, t)
	fmt.Println(s)
}
