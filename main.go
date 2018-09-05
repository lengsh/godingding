package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/lengsh/godingding/libs"
	"github.com/lengsh/godingding/mlog"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

func init() {
	// 设置默认数据库
}

type Msg struct {
	Message string
	Scrumb  string
}

func scrumbCreater(s string) string {
	secs := time.Now().Unix()
	pnum := secs / 60
	str := fmt.Sprintf("%s%d Scrumb secret keY", s, pnum)
	h := md5.New() // 计算MD5散列，引入crypto/md5 并使用 md5.New()方法。
	h.Write([]byte(str))
	bs := h.Sum(nil)
	log.Println(pnum)
	log.Println(fmt.Sprintf("%x", bs))
	return fmt.Sprintf("%x", bs)
}

func queryView(w http.ResponseWriter, r *http.Request) {
	var msg Msg
	t, _ := template.ParseFiles("view/query.gtpl")
	t.Execute(w, msg)
}

func firstPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	t, _ := template.ParseFiles("view/first.gtpl")
	var msg Msg
	t.Execute(w, msg)
}

func send(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		if m, ok := r.Form["message"]; ok {
			log.Println("send = ", m[0])
			if n, ok2 := r.Form[".scrumb"]; ok2 {
				scrum_new := scrumbCreater("send")
				if scrum_new == n[0] {
					dingtalker := libs.NewDingtalker()
					sm := template.HTMLEscapeString(m[0])
					dingtalker.SendChatTextMessage(sm)
				}
			}
		}
	}

	scr := scrumbCreater("send")
	var msg Msg = Msg{"", scr}
	t, _ := template.ParseFiles("view/send.gtpl")
	log.Println(t.Execute(w, msg))
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
					sm := template.HTMLEscapeString(m[0])
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
		log.Println(m)
		go func() {
			s := ""
			if content == nil {
				s = syscallDo("BABA")

			} else {
				s = syscallDo(strings.TrimSpace(fmt.Sprintf("%s", content)))
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

	sTime := libs.Crawler_163(msg)
	flog := mlog.LogInst()
	flog.LogInfo(sTime)
	sResult := libs.Crawler_Futu(msg)
	return fmt.Sprintf("%s\n %s", sTime, sResult)
}

func syscallDo(msg string) string {

	//  ProcMap := map[string]string{"text":"./bin/world","link":"./bin/stock"}
	sTime := libs.Crawler_163(msg)

	flog := mlog.LogInst()
	flog.LogInfo(sTime)
	/*
	   这里测试了一种通过启动了外部进程获得信息的方式，
	   如果后续有时间，可以增加通过plugin加载.so的方式。
	*/
	cmd := fmt.Sprintf("./bin/stock %s", msg)
	lsCmd := exec.Command("bash", "-c", cmd)
	lsOut, err := lsCmd.Output()
	if err != nil {
		// panic(err)
		flog.LogError(fmt.Sprintln(err))
		return "No Data!"
	} else {
		flog.LogInfo(string(lsOut))
		return fmt.Sprintf("%s\n %s", sTime, lsOut)
	}
}

func durationPing() {
	// 获得当前离明天早晨7点的时间距离, 即 每天早晨7点自动发送一条股市结果
	mt := time.Now().Unix()
	var ntt = 3600*24 - (mt%(3600*24) + 8*3600) + 7*3600
	var nt time.Duration = time.Duration(ntt)
	log.Printf("next report at ", ntt)
	time.AfterFunc(time.Duration(time.Second*nt), func() {
		// new log file mybe
		mlog.InitFilelog(true, "./log")
		s := syscallDo("BABA")
		dingtalker := libs.NewDingtalker()
		dingtalker.SendRobotTextMessage(s)
		durationPing()
	})

}

func main() {

	mlog.InitFilelog(true, "./log")
	flog := mlog.LogInst()
	flog.LogInfo("test mlog!!")

	durationPing()
	http.HandleFunc("/", firstPage)        //设置访问的路由
	http.HandleFunc("/send", send)         //设置访问的路由
	http.HandleFunc("/query", query)       //设置访问的路由
	http.HandleFunc("/help", help)         //设置访问的路由
	err := http.ListenAndServe(":80", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
