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
[root@yl-web yl]# yum install mariadb-server mariadb 
mariadb数据库的相关命令是：

systemctl start mariadb  #启动MariaDB
systemctl stop mariadb  #停止MariaDB
systemctl restart mariadb  #重启MariaDB
systemctl enable mariadb  #设置开机启动


utf8字集支持配置：
编辑 /etc/my.conf

init-connect='SET NAMES utf8'
# Settings user and group are ignored when systemd is used.
# If you need to run mysqld under a different user or group,
# customize your systemd unit file for mariadb according to the
# instructions in http://fedoraproject.org/wiki/Systemd
character-set-server = utf8
collation-server = utf8_unicode_ci

[mysqld_safe]
log-error=/var/log/mariadb/mariadb.log
pid-file=/var/run/mariadb/mariadb.pid
[mysqldump]
default-character-set=utf8
[mysql]
default-character-set=utf8


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
Install google-chrome & webdriver
配置yum下载源：
在目录 /etc/yum.repos.d/ 下新建文件 google-chrome.repo
[root@localhost ~]#    cd /ect/yum.repos.d/
[root@localhost yum.repos.d]#    vim google-chrome.repo
编辑google-chrome.repo，内容如下:
[google-chrome]
name=google-chrome
baseurl=http://dl.google.com/linux/chrome/rpm/stable/$basearch
enabled=1
gpgcheck=1
gpgkey=https://dl-ssl.google.com/linux/linux_signing_key.pub

安装google chrome浏览器：
 [root@localhost yum.repos.d]# yum -y install google-chrome-stable
 PS: Google官方源可能在中国无法使用，导致安装失败或者在国内无法更新，可以添加以下参数来安装：
 [root@localhost yum.repos.d]# yum -y install google-chrome-stable --nogpgcheck
 这样，google chrome就可在安装成功。

 安装chromedriver：

 [root@localhost ]#yum -y install chromedriver --nogpgcheck

 go lang代码跑起来：
 go get github.com/tebeka/selenium

 opts := []selenium.ServiceOption{}
 caps := selenium.Capabilities{
	      "browserName":                      "chrome",
  }
// 禁止加载图片，加快渲染速度
imagCaps := map[string]interface{}{
		  "profile.managed_default_content_settings.images": 2,
	  }

chromeCaps := chrome.Capabilities{
Prefs: imagCaps,
	       Path:  "",
	       Args: []string{
		       "--headless", // 设置Chrome无头模式
		       "--no-sandbox",
		       "--user-agent=Mozilla/5.0 (Linux; Android 8.1.0; EML-AL00 Build/HUAWEIEML-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Mobile Safari/537.36"},
	    }
caps.AddChrome(chromeCaps)
	// 启动chromedriver，端口号可自定义
	service, err := selenium.NewChromeDriverService("/usr/bin/chromedriver", 9515, opts...)
	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}

defer service.Stop()
	// 调起chrome浏览器
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		panic(err)
	}
defer webDriver.Quit()

	// 这是目标网站留下的坑，不加这个在linux系统中会显示手机网页，每个网站的策略不一样，需要区别处理。
	webDriver.AddCookie(&selenium.Cookie{
Name:  "defaultJumpDomain",
Value: "www",
})
// 导航到目标网站
urlBeijing := "http://m.iqiyi.com/vip/timeLine.html"
err = webDriver.Get(urlBeijing)
	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}
	fmt.Println(webDriver.Title())
str,err:= webDriver.PageSource()
	if err != nil {
		fmt.Println(err )
	}
fmt.Sprintf(str)
elem, err := webDriver.FindElement(selenium.ByClassName, "m-vip-timer-shaft")  //ByCSSSelector, "m-vip-timer-shaft")
if err != nil {
	panic(err)
}

##获取一个a标签下的URL连接地址：

melem, err := r.webDriver.FindElement(selenium.ByClassName, "findSection") //article")
if err != nil {
	logs.Error("findSection >>>> ", err)
		return ""
}

ss, err := melem.Text()
	if err != nil {
		logs.Error(err)
	} else {
svs := strings.Split(ss, "\n")
	     if len(svs) > 1 {
s := svs[1]
	   if svs[1] == "Titles" {
		   s = svs[2]
	   }
svsv := strings.Split(s, "(")
	      s1 := svsv[0]
	      mv = strings.TrimSpace(s1)
	     }
	}
mele, err := melem.FindElement(selenium.ByLinkText, mv)
	if err != nil {
		logs.Error("link Text >>>> ", err)
			return ""
	}

aurl, err := mele.GetAttribute("href")
if err != nil {
	logs.Error(err)
		return ""
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

var corpid string = "ding5b26ca................f4657eb6378f"
var corpsecret string = "2uK2a27AWgkfk.......................boCivL_Gry"

func main() {
   c := godingtalk.NewDingTalkClient(corpid, corpsecret)
   c.RefreshAccessToken()
   var chatid string = "chat8890dbc..................85e"
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
# 6. 微信消息和自定义机器人 
```
https://www.jianshu.com/p/96a969ad4b02

## 后台执行服务 
nohup让提交的命令忽略 hangup 信号( run a command immune to hangups, with output to a non-tty)
sudo nohup godingding 2>&1 &
or
sudo setsid godingding &


[参见]
 https://www.ibm.com/developerworks/cn/linux/l-cn-nohup/index.html
 https://github.com/sevlyar/go-daemon

## vim配置go语法高亮
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
