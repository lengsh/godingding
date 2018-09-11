package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/jasonlvhit/gocron"
	"github.com/lengsh/godingding/libs"
	// "github.com/lengsh/godingding/log4go"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

func init() {

}

type Msg struct {
	Message string
	Scrumb  string
}

func main() {

	logs.SetLogger(logs.AdapterFile, `{"filename":"./log/godingding.log","maxlines":10000,"maxsize":102400,"daily":true,"maxdays":2}`)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)

	http.HandleFunc("/", firstPage)  //设置访问的路由
	http.HandleFunc("/send", send)   //设置访问的路由
	http.HandleFunc("/query", query) //设置访问的路由
	http.HandleFunc("/help", help)   //设置访问的路由

	scheduler := gocron.NewScheduler()
	job1 := scheduler.Every(1).Day().At("07:01")
	job1.Do(durationPing) // durationPing()
	scheduler.Start()

	err := http.ListenAndServe(":80", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func scrumbCreater(s string) string {
	secs := time.Now().Unix()
	pnum := secs / 60
	str := fmt.Sprintf("%s%d Scrumb secret keY", s, pnum)
	h := md5.New() // 计算MD5散列，引入crypto/md5 并使用 md5.New()方法。
	h.Write([]byte(str))
	bs := h.Sum(nil)
	//	log.Println(pnum)
	//	log.Println(fmt.Sprintf("%x", bs))
	return fmt.Sprintf("%x", bs)
}

func queryView(w http.ResponseWriter, r *http.Request) {
	qs := libs.QueryStock()
	t, _ := template.ParseFiles("view/query.gtpl")
	err := t.Execute(w, qs)
	if err != nil {
		logs.Error(fmt.Sprint(err))
	}
}

func firstPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据

	qs := libs.QueryStock()
	if qs != nil {
		t, _ := template.ParseFiles("view/query.gtpl")
		err := t.Execute(w, qs)
		if err != nil {
			logs.Error(fmt.Sprint(err))
		}
	} else {
		t, _ := template.ParseFiles("view/first.gtpl")
		err := t.Execute(w, nil)
		if err != nil {
			logs.Error(fmt.Sprint(err))
		}
	}
}

func send(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		r.ParseForm()
		if m, ok := r.Form["message"]; ok {
			if n, ok2 := r.Form[".scrumb"]; ok2 {
				scrum_new := scrumbCreater("send")
				if scrum_new == n[0] {
					dingtalker := libs.NewDingtalker()
					sm := m[0] // template.HTMLEscapeString(m[0])
					dingtalker.SendChatTextMessage(sm)
				}
			}
		}
	}
	scr := scrumbCreater("send")
	var msg Msg = Msg{"", scr}
	t, _ := template.ParseFiles("view/send.gtpl")
	err := t.Execute(w, msg)
	if err != nil {
		logs.Error(fmt.Sprint(err))
	}
}

func query(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//	queryView(w, r)
		r.ParseForm()
		if m, ok := r.Form["message"]; ok {
			if n, ok2 := r.Form[".scrumb"]; ok2 {
				scrum_new := scrumbCreater("send")
				if scrum_new == n[0] {
					dingtalker := libs.NewDingtalker()
					sm := m[0] //  template.HTMLEscapeString(m[0])
					dingtalker.SendRobotTextMessage(sm)
				}
			}
		}

	} else {
	}
	firstPage(w, r)
}

func help(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()

		var f interface{}
		json.Unmarshal(result, &f)
		m := f.(map[string]interface{})
		senderNick := m["senderNick"]
		text1 := m["text"]
		text2 := text1.(map[string]interface{})
		content := text2["content"]
		go func() {
			s := ""
			if content == nil {
				s = syscallDo("BABA")

			} else {
				//sm := template.HTMLEscapeString(strings.TrimSpace(fmt.Sprintf("%s", content)))
				sm := strings.TrimSpace(fmt.Sprintf("%s", content))
				//s = syscallDo(sm)
				s = pluginDo(sm)
			}
			ss := fmt.Sprintf("@%s\n%s", senderNick, s)
			dingtalker := libs.NewDingtalker()
			dingtalker.SendRobotTextMessage(ss)

		}()
	}
}

/*

 通过plugin加载.so的方式。

*/
func pluginDo(msg string) string {

	so := libs.Plugins{"./so/stockplugin.so"}
	sTime := so.Crawler_Stock(msg)
	logs.Info("Stock Time = ", sTime)
	st := libs.Crawler_Futu(msg, sTime)
	st.NewStock()
	return st.String()
}

func syscallDo(msg string) string {

	//  ProcMap := map[string]string{"text":"./bin/world","link":"./bin/stock"}
	sTime := libs.Crawler_163(msg)
	logs.Info(sTime)
	/*
	   这里测试了一种通过启动了外部进程获得信息的方式，
	   如果后续有时间，可以增加通过plugin加载.so的方式。
	*/
	cmd := fmt.Sprintf("./bin/stock %s", msg)
	lsCmd := exec.Command("bash", "-c", cmd)
	lsOut, err := lsCmd.Output()
	if err != nil {
		// panic(err)
		logs.Error(fmt.Sprintln(err))
		return "No Data!"
	} else {
		logs.Info(string(lsOut))
		return fmt.Sprintf("%s\n %s", sTime, lsOut)
	}
}

func durationPing() {
	// 获得当前离明天早晨7点的时间距离, 即 每天早晨7点自动发送一条股市结果
	sk := "BABA"
	s := pluginDo(sk)
	dingtalker := libs.NewDingtalker()
	dingtalker.SendRobotTextMessage(s)
}
