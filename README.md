# godingding
```
```
# 1. 了解 web server、View模版的知识
````
  var stk []Stockorm
  stk := GetStocksFromDB() //TODO
  t, _ := template.ParseFiles("query.gtpl")
  t.Execute(w, stk)

-----------  query.gtpl ----------
 {{range .}}
    <tr><th>{{.Id}} </th><th>{{.Name}} </th><th> {{.Department}} </th><th>{{.Number}}</th></tr>
 {{end}}

https://beego.me/docs/intro/

````
# 2. 动态库生成和调用plugin
```
go build -buildmode=plugin stockplugin.go 

----- stockplugin.so
 func CrawlerStock(st string) string {
         return libs.Crawler_163(st)
 }
-------
 p,_ := plugin.Open("./so/stockplugin.so")
 crawlerstock, err := p.Lookup("CrawlerStock")
 result = crawlerstock.(func(string) string)(stk)

```
# 3. 数据库及ORM
````
type Stock struct {
	Name        string  `orm:"size(20); index"`
	HighPrice   float64 `orm: "default(0)"`
	TradeDate   string `orm:"size(32); index"`
}

type Stockorm struct {
	Id int
	Stock
}
func init() {
	// 设置默认数据库
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	// 设置默认数据库，数据库存放位置：./datas/test.db ， 数据库别名：default
	orm.RegisterDataBase("default", "sqlite3", "./test.db")
	// 注册定义的 model
	orm.RegisterModel(new(Stockorm))
	orm.RunSyncdb("default", false, true)
}

o := orm.NewOrm()
var rs orm.RawSeter
sql := fmt.Sprintf("SELECT * FROM stockorm WHERE  name ='%s' AND trade_date ='%s'", r.Name, r.TradeDate)
rs = o.Raw(sql)
var stocks []Stockorm
num, err := rs.QueryRows(&stocks)

https://beego.me/docs/mvc/model/overview.md
https://astaxie.gitbooks.io/build-web-application-with-golang/zh/05.5.html


````
# 4. 动态网页爬取，https
```
下载phantomjs并解码，设置运行路径。
http://phantomjs.org/download.html

PATH=/Users/lengss/devel/phantomjs-2.1.1-maco/bin:${PATH}"
. ~/.bash_profile
go get -u github.com/benbjohnson/phantomjs

import "github.com/benbjohnson/phantomjs"
import "testing"
import "fmt"
import "os"

func Test1_abc(t *testing.T) {
	// Start the process once.
	if err := phantomjs.DefaultProcess.Open(); err != nil {
		fmt.Println(err)
			os.Exit(1)
	}
	defer phantomjs.DefaultProcess.Close()
		page,err := phantomjs.CreateWebPage()
		if err != nil {
			return
		}

	if err := page.Open("http://quotes.money.163.com/usstock/BABA.html#US1a01"); err != nil {
		return 
	}

	// Setup the viewport and render the results view.
	if err := page.SetViewportSize(1024, 800); err != nil {
		return 
	}
	if content,err := page.Content();err == nil{
		fmt.Printf(content)
	}
}


https://github.com/benbjohnson/phantomjs/blob/master/README.md
https://github.com/henrylee2cn/pholcus

```
# 5. dingtalk的消息发送与webhook（机器人）
````

申请获取corpID和corpsecret：
https://oa.dingtalk.com
https://open-dev.dingtalk.com/#/corpAuthInfo
[userid]: 企业通讯录中可以获得userid
[AgentID]: 微应用申请分配的ID

go get github.com/hugozhu/godingtalk

code example:

package main
import (
		"github.com/hugozhu/godingtalk"
		"fmt"
       )

var corpid string = "ding5b26ca68f242cff035c2f4657eb6378f"
var corpsecret string = "2uK2a27AWgkfkVAxd9IdwqG9SO7D01LhWnCgDEYhxff6uGj924NEdrboCivL_Gry"

func main() {
   c := godingtalk.NewDingTalkClient(corpid, corpsecret)
   c.RefreshAccessToken()
   var chatid string = "chat8890dbc9d98595c5a1031fe99d8c585e"
   err = c.SendTextMessage("bc0002", chatid, "测试消息，请忽略3")
   if err != nil {
	   fmt.Println(err)
   }
}

在钉钉群里引入自定义机器人，并设置好webhook：

  c := godingtalk.NewDingTalkClient(r.CorpId, r.CorpSecret)
  if c != nil {
	c.RefreshAccessToken()
	err := c.SendRobotTextMessage(r.AcToken, msg)
	if err != nil {
			log.Println(err)
	}
  }

https://github.com/hugozhu/godingtalk

````
# 6. 
```

```
==========================
```
nohup让提交的命令忽略 hangup 信号( run a command immune to hangups, with output to a non-tty)
sudo nohup godingding &
or
sudo setsid godingding &

[参见]
 https://www.ibm.com/developerworks/cn/linux/l-cn-nohup/index.html
 https://github.com/sevlyar/go-daemon

```
==========================
```
vim配置go语法高亮
cd ~
mkdir .vim
cd .vim
mkdir autoload  plugged
cd plugged
git clone https://github.com/fatih/vim-go vim-go
cd ../autoload
wget https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
配置vimrc文件：
[root@localhost ~]#vim ~/.vimrc
增加：
call plug#begin()
Plug 'fatih/vim-go', { 'do': ':GoInstallBinaries' }
call plug#end()
let g:go_version_warning = 0
```

