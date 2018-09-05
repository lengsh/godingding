package main

import (
	"fmt"
	"github.com/lengsh/godingding/libs"
	"os"
	"strings"
)

func main() {
	stk := "BABA"
	if len(os.Args) > 1 {
		stk = os.Args[1]
		stk = strings.ToUpper(stk)
	}
	s := libs.Crawler_Futu(stk)
	fmt.Println(s)

}
