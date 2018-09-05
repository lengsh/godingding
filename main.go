package main

import (
	"fmt"
"strings"
"html/template"
	"log"
        "io/ioutil"
        "net/http"
        "os/exec"
        "encoding/json"
 "crypto/md5"
"time"
"github.com/lengsh/godingding/libs"
"github.com/lengsh/godingding/mlog"

)

func init() {
	// 设置默认数据库
}

type Msg struct {
	Message string  
        Scrumb string
}

func scrumbCreater(s string) string{
    secs := time.Now().Unix()
    pnum := secs/60
    str := fmt.Sprintf("%s%d Scrumb secret keY",s, pnum)
    h := md5.New()   // 计算MD5散列，引入crypto/md5 并使用 md5.New()方法。
    h.Write([]byte(str))
    bs := h.Sum(nil)
    log.Println(pnum)
    log.Println(fmt.Sprintf("%x",bs))
   return fmt.Sprintf("%x",bs)
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
                 dingtalker.SendChatTextMessage(m[0])
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
                         dingtalker.SendRobotTextMessage(m[0])
	           }
		}
		}

	} else {
	}
	firstPage(w,r)
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
log.Println( m )
       go func(){
             s:=""
             if(content == nil){
               s = syscallDo("BABA")

           } else {
               s = syscallDo(strings.TrimSpace(fmt.Sprintf("%s",content)))
            }
           ss := fmt.Sprintf("@%s\n%s", senderNick , s)
	   dingtalker := libs.NewDingtalker()
           dingtalker.SendRobotTextMessage( ss )
           
       }()
   }
}


func syscallDo(msg string) string  {

    //  ProcMap := map[string]string{"text":"./bin/world","link":"./bin/stock"}

     flog := mlog.LogInst()
     flog.LogInfo("test mlog!!")

     cmd := fmt.Sprintf("./bin/stock %s", msg)
     lsCmd := exec.Command("bash", "-c", cmd)
     lsOut, err := lsCmd.Output()
     if err != nil {
	        // panic(err)
          flog.LogError(fmt.Sprintln(err))
	  return "No Data!"
     } else{
        flog.LogInfo(string(lsOut))
        return  fmt.Sprintf("%s", lsOut)
     }
}

 func durationPing(){
	time.AfterFunc(time.Duration(3600*time.Second), func() {
s := syscallDo("BABA")
dingtalker := libs.NewDingtalker()
dingtalker.SendRobotTextMessage( s )
    durationPing()
})
}

func main() {

   mlog.InitFilelog(true, "./log")
   flog := mlog.LogInst()
   flog.LogInfo("test mlog!!")
  
   durationPing()
http.HandleFunc("/", firstPage)  //设置访问的路由
	http.HandleFunc("/send", send)     //设置访问的路由
	http.HandleFunc("/query", query) //设置访问的路由
	http.HandleFunc("/help", help) //设置访问的路由
	err := http.ListenAndServe(":80", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
