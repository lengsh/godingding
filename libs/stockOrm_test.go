package libs

import (
	"fmt"
	"github.com/lengsh/godingding/libs"
	"testing"
)

func TestNewStock(t *testing.T) {
	st := "2018-09-04 19:45:32"
	stock := libs.Stock{Name: "BABA", HighPrice: 20, LowPrice: 10, StartPrice: 11, EndPrice: 15, TradeSum: 120, TradeVol: 200.12, TradeDate: st}
	n := stock.NewStock()
	if n == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("error to insert :", n)
	}
}
