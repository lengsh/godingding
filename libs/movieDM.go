package libs

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
)

// Model Struct
// https://beego.me/docs/mvc/model/models.md#%E6%A8%A1%E5%9E%8B%E5%AD%97%E6%AE%B5%E4%B8%8E%E6%95%B0%E6%8D%AE%E5%BA%93%E7%B1%BB%E5%9E%8B%E7%9A%84%E5%AF%B9%E5%BA%94

type TagMovie struct {
	Movie
	TagRate string
}

func (r Movie) NewMovie() int {
	o := orm.NewOrm()
	var rs orm.RawSeter
	sql := fmt.Sprintf("SELECT * FROM movie WHERE  company ='%s' AND  name ='%s'", r.Company, r.Name)
	logs.Debug(sql)
	rs = o.Raw(sql)
	var ms []Movie
	num, err := rs.QueryRows(&ms)
	if err != nil {
		logs.Error(err)
	} else if num < 1 {
		//		var ns Stockorm = Stockorm{0, r}
		id, err := o.Insert(&r)
		if err != nil {
		} else {
			logs.Info(id)
			return 1
		}
	} else {
		logs.Debug("return count must be 1 === ", num)
		r.Id = ms[0].Id
		id, err := o.Update(&r)
		if err != nil {
		} else {
			logs.Info(id)
			return 1
		}
	}
	return 0
}
func QueryLastMovies(max int) []Movie {
	o := orm.NewOrm()
	var rs orm.RawSeter
	if max > 100 {
		max = 100
	}
	sql := fmt.Sprintf("SELECT * FROM movie ORDER BY releasetime desc LIMIT %d", max)
	logs.Debug(sql)
	rs = o.Raw(sql)
	var ms []Movie
	_, err := rs.QueryRows(&ms)
	if err != nil {
		logs.Error(err)
		return nil
	} else {
		return ms
	}
}

func QueryCompanyMovies(com string, max int) []Movie {
	if max > 100 {
		logs.Error("num is too big!!")
		return nil
	}

	o := orm.NewOrm()
	var rs orm.RawSeter
	sql := fmt.Sprintf("SELECT * FROM movie WHERE company='%s' ORDER BY releasetime desc LIMIT %d", com, max)
	logs.Debug(sql)
	rs = o.Raw(sql)
	var ms []Movie
	_, err := rs.QueryRows(&ms)
	if err != nil {
		logs.Error(err)
		return nil
	} else {
		return ms
	}
}

func QueryNameMovies(name string) []Movie {
	o := orm.NewOrm()
	var rs orm.RawSeter
	sql := fmt.Sprintf("SELECT * FROM movie WHERE name='%s'", name)
	logs.Debug(sql)
	rs = o.Raw(sql)
	var ms []Movie
	_, err := rs.QueryRows(&ms)
	if err != nil {
		logs.Error(err)
		return nil
	} else {
		return ms
	}
}

func (r Movie) String() string {
	return fmt.Sprintln("Company：", r.Company, "\nName：", r.Name, "\nRate：", r.Rate, "\nRelease Time：", r.Releasetime)
}
