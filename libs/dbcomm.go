package libs

import (
	//	"fmt"
	//	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
)

// Model Struct
// https://beego.me/docs/mvc/model/models.md#%E6%A8%A1%E5%9E%8B%E5%AD%97%E6%AE%B5%E4%B8%8E%E6%95%B0%E6%8D%AE%E5%BA%93%E7%B1%BB%E5%9E%8B%E7%9A%84%E5%AF%B9%E5%BA%94

type Stock struct {
	Name        string  `orm:"size(20); index"`
	HighPrice   float64 `orm: "default(0)"`
	LowPrice    float64
	StartPrice  float64
	EndPrice    float64
	TradeStock  float64
	TradeFounds float64
	TradeDate   string `orm:"size(32); index"`
}

type Stockorm struct {
	Id int
	Stock
}

type Movie struct {
	Id          int
	Company     string  `orm:"size(20);index"`
	Name        string  `orm:"size(32);index"`
	Rate        float32 `orm: "default(0)"`
	Releasetime string  `orm:"size(32); index"`
}

func init() {
	// 设置默认数据库
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "lengsh:hsgnel@/youku?charset=utf8")
	//("mysql", "user:password@/dbname")
	// 注册定义的 model

	orm.RegisterModel(new(Movie), new(Stockorm))
	orm.RunSyncdb("default", false, true)

}
