package libs

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
	//	"github.com/lengsh/godingding/log4go"
	//_ "github.com/mattn/go-sqlite3"
	//     _ "github.com/go-sql-driver/mysql" // 导入数据库驱动
)

// Model Struct
// https://beego.me/docs/mvc/model/models.md#%E6%A8%A1%E5%9E%8B%E5%AD%97%E6%AE%B5%E4%B8%8E%E6%95%B0%E6%8D%AE%E5%BA%93%E7%B1%BB%E5%9E%8B%E7%9A%84%E5%AF%B9%E5%BA%94
type StockSum struct {
	SumMarket float32
	Date      time.Time
}

func (r Stock) NewStock() int {
	o := orm.NewOrm()
	var rs orm.RawSeter
	s := fmt.Sprintf("%d-%02d-%02d", r.CreateDate.Year(), r.CreateDate.Month(), r.CreateDate.Day())
	sql := fmt.Sprintf("SELECT * FROM stockorm WHERE  name ='%s' AND create_date ='%s'", r.Name, s)
	logs.Debug(sql)

	rs = o.Raw(sql)
	var stocks []Stockorm
	num, err := rs.QueryRows(&stocks)
	if err != nil {
		logs.Error(err)
	} else if num < 1 {
		var ns Stockorm = Stockorm{0, r}
		id, err := o.Insert(&ns)
		if err != nil {
		} else {
			logs.Info(id)
			return 1
		}
	}
	logs.Debug("data is exist!")
	return 0
}

func (r Stock) SumMarketCap() float32 {
	s := fmt.Sprintf("%d-%02d-%02d", r.CreateDate.Year(), r.CreateDate.Month(), r.CreateDate.Day())
	return QueryMarketCap(s)
}

func QueryStock(num int) []Stockorm {
	o := orm.NewOrm()
	var rs orm.RawSeter
	sql := fmt.Sprintf("SELECT * FROM stockorm ORDER BY create_date desc LIMIT %d", num)
	logs.Debug(sql)
	rs = o.Raw(sql)
	var stocks []Stockorm
	_, err := rs.QueryRows(&stocks)
	if err != nil {
		logs.Error(err)
		return nil
	} else {
		return stocks
	}
}

func QueryMarketCaps(dt string) []StockSum {
	o := orm.NewOrm()
	var rs orm.RawSeter
	sql := fmt.Sprintf("SELECT  sum(market_cap) as sum_market, create_date as date from stockorm where create_date >='%s' group by create_date", dt)
	logs.Debug(sql)
	rs = o.Raw(sql)

	var sa []StockSum
	_, err := rs.QueryRows(&sa)
	if err != nil {
		logs.Error(err)
		return nil
	} else {
		return sa
	}
}

func QueryMarketCap(dt string) float32 {
	o := orm.NewOrm()
	var rs orm.RawSeter
	sql := fmt.Sprintf("SELECT  sum(market_cap) as sum_market, create_date from stockorm where create_date='%s' group by create_date", dt)
	logs.Debug(sql)
	rs = o.Raw(sql)

	var sa StockSum
	err := rs.QueryRow(&sa)
	if err != nil {
		logs.Error(err)
		return 0
	} else {
		return sa.SumMarket
	}
}

//  select sum(market_cap), create_date from stockorm group by create_date;

func QueryOneStock(sID string, start int, num int) []Stockorm {
	stk := strings.ToUpper(sID)
	o := orm.NewOrm()
	var rs orm.RawSeter
	sql := fmt.Sprintf("SELECT * FROM stockorm WHERE name='%s' ORDER BY create_date desc LIMIT %d,%d", stk, start, num)
	logs.Debug(sql)
	rs = o.Raw(sql)
	var stocks []Stockorm
	_, err := rs.QueryRows(&stocks)
	if err != nil {
		logs.Error(err)
		return nil
	} else {
		return stocks
	}
}

func LastStock(stock string) (Stockorm, error) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	s := strings.ToUpper(stock)
	sql := fmt.Sprintf("SELECT * FROM stockorm WHERE name='%s' order by create_date desc LIMIT 1", s)
	logs.Debug(sql)
	rs = o.Raw(sql)
	var stocks []Stockorm
	_, err := rs.QueryRows(&stocks)
	if err != nil {
		logs.Error(err)
		return Stockorm{}, err
	} else {
		if len(stocks) == 1 {
			return stocks[0], nil
		} else {
			return Stockorm{}, errors.New("no data")
		}
	}
}

func (r Stock) TString() string {
	return fmt.Sprintf("%d-%02d-%02d", r.CreateDate.Year(), r.CreateDate.Month(), r.CreateDate.Day())
}
func (r StockSum) TString() string {
	return fmt.Sprintf("%d-%02d-%02d", r.Date.Year(), r.Date.Month(), r.Date.Day())
}
func (r Stock) String() string {
	t := fmt.Sprintf("%d年%02d月%02d日", r.CreateDate.Year(), r.CreateDate.Month(), r.CreateDate.Day())
	return fmt.Sprintln("代码：", r.Name, "\n时间：", t, "\n最高价：", r.HighPrice, "\n最低价：", r.LowPrice, "\n开盘价：", r.StartPrice, "\n当前价：", r.EndPrice, "\n成交额：", r.TradeFounds, "亿\n成交量：", r.TradeStock, "万\n市值：", r.MarketCap, "亿")

}
