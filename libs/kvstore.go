package libs

import (
	//	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	//	"math"
	//	"strings"
	"time"
	//	"github.com/lengsh/godingding/log4go"
	//_ "github.com/mattn/go-sqlite3"
	//     _ "github.com/go-sql-driver/mysql" // 导入数据库驱动
)

// Model Struct
// https://beego.me/docs/mvc/model/models.md#%E6%A8%A1%E5%9E%8B%E5%AD%97%E6%AE%B5%E4%B8%8E%E6%95%B0%E6%8D%AE%E5%BA%93%E7%B1%BB%E5%9E%8B%E7%9A%84%E5%AF%B9%E5%BA%94

func GetKVStore(domain string, key string) (string, bool) {
	//	t := time.Now().UTC().Add(8 * time.Hour)
	o := orm.NewOrm()
	var rs orm.RawSeter
	sql := fmt.Sprintf("SELECT * FROM kvstore where domain='%s' AND kkey='%s'", domain, key)
	logs.Debug(sql)
	rs = o.Raw(sql)
	var sk []Kvstore
	num, err := rs.QueryRows(&sk)

	if err != nil || num < 1 {
		logs.Error(err)
		return "", false
	} else {
		return sk[0].Vvalue, true
	}
}

func (r TouTiao) UpdateTime() string {
	key := time.Now().Format("2006-01-02")
	if t, b := GetKVStoreTime("RESOU", key); b {
		return t.Format("2006-01-02 15:04:05")
	}
	return ""
}

func GetKVStoreTime(domain string, key string) (time.Time, bool) {
	o := orm.NewOrm()
	var rs orm.RawSeter
	sql := fmt.Sprintf("SELECT * FROM kvstore where domain='%s' AND kkey='%s'", domain, key)
	logs.Debug(sql)
	rs = o.Raw(sql)
	var sk []Kvstore
	num, err := rs.QueryRows(&sk)

	if err != nil || num < 1 {
		logs.Error(err)
		t := time.Now().UTC().Add(8 * time.Hour)
		return t, false
	} else {
		return sk[0].Modetime, true
	}
}

func SetKVStore(domain string, key string, value string) bool {
	o := orm.NewOrm()
	var rs orm.RawSeter
	sql := fmt.Sprintf("SELECT * FROM kvstore where domain='%s' AND kkey='%s'", domain, key)
	logs.Debug(sql)
	rs = o.Raw(sql)
	var sk []Kvstore
	num, _ := rs.QueryRows(&sk)
	if num >= 1 {
		sk[0].Vvalue = value
		sk[0].Modetime = time.Now().UTC() // .Add(8 * time.Hour)

		if _, err := o.Update(&sk[0]); err == nil {
			return true
		} else {
			logs.Debug(err)
			return false
		}
	} else {
		var ns Kvstore = Kvstore{0, domain, key, value, time.Now().UTC()}
		id, err := o.Insert(&ns)
		if err != nil {
		} else {
			logs.Info(id)
			return false
		}

	}
	return true
}
