package libs

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	//"time"
	//     _ "github.com/go-sql-driver/mysql" // 导入数据库驱动
)

// Model Struct
// https://beego.me/docs/mvc/model/models.md#%E6%A8%A1%E5%9E%8B%E5%AD%97%E6%AE%B5%E4%B8%8E%E6%95%B0%E6%8D%AE%E5%BA%93%E7%B1%BB%E5%9E%8B%E7%9A%84%E5%AF%B9%E5%BA%94

type Stock struct {
	Id         int
	Name       string  `orm:"size(20); index"`
	HighPrice  float64 `orm: "default(0)"`
	LowPrice   float64
	StartPrice float64
	EndPrice   float64
	TradeSum   float64
	TradeVol   float64
	TradeDate  string `orm:"size(32); index"`
}

// Id, HighPrice, LowPrice, StarPrice, EndPrice, TradSum, TradVol, TradDate

func init() {
	// 设置默认数据库
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	// 设置默认数据库，数据库存放位置：./datas/test.db ， 数据库别名：default
	orm.RegisterDataBase("default", "sqlite3", "./test.db")
	// 注册定义的 model
	orm.RegisterModel(new(Stock))
	orm.RunSyncdb("default", false, true)

}

func (r Stock) NewStock() int {
	o := orm.NewOrm()
	var rs orm.RawSeter
	sql := fmt.Sprintf("SELECT * FROM stock WHERE  name ='%s' AND trade_date ='%s'", r.Name, r.TradeDate)
	fmt.Println(sql)
	rs = o.Raw(sql)
	var stocks []Stock
	num, err := rs.QueryRows(&stocks)
	if err != nil {
		fmt.Println(err)
	} else if num < 1 {
		id, err := o.Insert(&r)
		if err != nil {
		} else {
			fmt.Println(id)
			return 1
		}
	}
	return 0
}
