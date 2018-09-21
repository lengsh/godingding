package libs

import (
	"fmt"
	"github.com/lengsh/godingding/libs"
	// "os"
	"testing"
)

func init() {
}

func TestNewStock(t *testing.T) {
	st := "2018-09-05 19:25:32"
	stock := libs.Stock{Name: "BABA", HighPrice: 20, LowPrice: 10, StartPrice: 11, EndPrice: 15, TradeFounds: 120, TradeStock: 200.12, TradeDate: st}
	n := stock.NewStock()
	if n == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("error to insert :", n)
	}
}

func TestQueryStock(t *testing.T) {
	qs := libs.QueryStock()
	if qs != nil {
		for _, s := range qs {
			fmt.Println("Id = ", s.Id)
			fmt.Printf(s.String())
		}
	}

}
