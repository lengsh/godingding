package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/jasonlvhit/gocron"
	"github.com/lengsh/godingding/libs"
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
	http.HandleFunc("/stock", stock) //设置访问的路由
	http.HandleFunc("/help", help)   //设置访问的路由

	scheduler := gocron.NewScheduler()
	job1 := scheduler.Every(1).Day().At("07:01")
	job1.Do(stockPing)

	job2 := scheduler.Every(1).Day().At("10:49")
	job2.Do(crawMovieJob)

	job3 := scheduler.Every(1).Day().At("20:01")
	job3.Do(crawMovieJob)
	//scheduler.Every(1).Day().Do(crawJob)
	// scheduler.Every(1).Hour().Do(crawJob)
	scheduler.Start()

	err := http.ListenAndServe(":8081", nil) //设置监听的端口
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

func queryStock(w http.ResponseWriter, r *http.Request) {
	qs := libs.QueryStock(10)
	t, _ := template.ParseFiles("view/stock.gtpl", "view/header.gtpl", "view/footer.gtpl")
	err := t.Execute(w, qs)
	if err != nil {
		logs.Error(err.Error())
	}
}

func movieReport(w http.ResponseWriter, r *http.Request) {
	tx := libs.QueryTopMovies("TX", 100)
	yk := libs.QueryTopMovies("YOUKU", 100)
	qy := libs.QueryTopMovies("IQIYI", 100)

	var qs []libs.Movie
	qs = append(qs, tx...)
	qs = append(qs, yk...)
	qs = append(qs, qy...)

	t, _ := template.ParseFiles("view/report.gtpl", "view/tx.gtpl", "view/iqiyi.gtpl", "view/youku.gtpl")
	err := t.Execute(w, qs)
	if err != nil {
		logs.Debug(err.Error())
		logs.Error(err.Error())
	}
}

func queryMovie(w http.ResponseWriter, r *http.Request) {

	tx := libs.QueryXMovies("TX", 100)
	yk := libs.QueryXMovies("YOUKU", 100)
	qy := libs.QueryXMovies("IQIYI", 100)

	var qs []libs.Movie
	qs = append(qs, tx...)
	qs = append(qs, yk...)
	qs = append(qs, qy...)

	t, _ := template.ParseFiles("view/movie.gtpl", "view/header.gtpl", "view/footer.gtpl")
	err := t.Execute(w, qs)
	if err != nil {
		logs.Debug(err.Error())
		logs.Error(err.Error())
	}
}

func firstPage(w http.ResponseWriter, r *http.Request) {

	qs1 := libs.QueryTopMovies("IQIYI", 10)
	qs2 := libs.QueryTopMovies("TX", 10)
	qs3 := libs.QueryTopMovies("YOUKU", 10)
	var qs []libs.Movie

	qs = append(qs, qs1...)
	qs = append(qs, qs2...)
	qs = append(qs, qs3...)

	t, _ := template.ParseFiles("view/first.gtpl", "view/header.gtpl", "view/footer.gtpl")
	err := t.Execute(w, qs)
	if err != nil {
		logs.Error(err.Error())
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
	t, _ := template.ParseFiles("view/send.gtpl", "view/header.gtpl", "view/footer.gtpl")
	err := t.Execute(w, msg)
	if err != nil {
		logs.Error(err.Error())
	}
}

func query(w http.ResponseWriter, r *http.Request) {
	key := 0
	if r.Method == "GET" {
		//	queryView(w, r)
		r.ParseForm()
		if m, ok := r.Form["do"]; ok {
			if m[0] == "stock" {
				key = 1
			} else if m[0] == "movie" {
				key = 2
			} else if m[0] == "report" {
				key = 3
			}
		}
	}
	switch key {
	case 3:
		movieReport(w, r)
	case 2:
		queryMovie(w, r)
	case 1:
		queryStock(w, r)
	default:
		t, _ := template.ParseFiles("view/first.gtpl", "view/header.gtpl", "view/footer.gtpl")
		t.Execute(w, nil)
	}
}

func stock(w http.ResponseWriter, r *http.Request) {
	key := "baba"
	if r.Method == "GET" {
		r.ParseForm()
		if m, ok := r.Form["do"]; ok {
			key = m[0]
			sk := libs.QueryOneStock(key, 0, 100)
			t, _ := template.ParseFiles("view/onestock.gtpl", "view/header.gtpl", "view/footer.gtpl")
			err := t.Execute(w, sk)
			if err != nil {
				logs.Error(err.Error())
			}
		} else {
			tn := time.Now().UTC().Add(8 * time.Hour)
			tt := tn.Add(-100 * 24 * time.Hour)
			s := fmt.Sprintf("%d-%2d-%2d", tt.Year(), tt.Month(), tt.Day())
			sk := libs.QueryMarketCaps(s)
			t, _ := template.ParseFiles("view/stocksum.gtpl", "view/header.gtpl", "view/footer.gtpl")
			err := t.Execute(w, sk)
			if err != nil {
				logs.Error(err.Error())
			}
		}
	}
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
		keyword := strings.TrimSpace(fmt.Sprintf("%s", content))
		keyword = strings.ToUpper(keyword)
		go func() {
			switch keyword {
			case "XIANXING":
				fallthrough
			case "限行":
				fallthrough
			case "CAR LIMIT":
				s := libs.CrawlCarLimitJob()
				ss := fmt.Sprintf("@%s\n%s", senderNick, s)
				dingtalker := libs.NewDingtalker()
				dingtalker.SendRobotTextMessage(ss)
			case "爬股票":
				fallthrough
			case "UPDATE STOCK":
				libs.CrawlStocksJob()
			case "爬电影":
				fallthrough
			case "UPDATE MOVIE":
				libs.CrawlMovieJob()
			default:
				stocks := "BABA,FB,MSFT,AMZN,AAPL,TSLA,BIDU,NVDA,GOOGL,WB"
				s := ""
				if strings.Contains(stocks, keyword) {
					s = libs.CrawlStockJob(keyword)
				} else {
					s = pluginDo(keyword)
				}
				ss := fmt.Sprintf("@%s\n%s", senderNick, s)
				dingtalker := libs.NewDingtalker()
				dingtalker.SendRobotTextMessage(ss)
			}
		}()
	}

	if r.Method == "GET" {
		// return data/xxx.data
	}
}

/*

 通过plugin加载.so的方式。

*/
func pluginDo(msg string) string {
	/*
		so := libs.Plugins{"./so/stockplugin.so"}
		sTime := so.Crawler_Stock(msg)
		logs.Info("Stock Time = ", sTime)
	*/
	st := libs.Crawler_Futu(msg)
	//st.NewStock()
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
		logs.Error(err.Error())
		return "No Data!"
	} else {
		logs.Info(string(lsOut))
		return fmt.Sprintf("%s\n %s", sTime, lsOut)
	}
}

func stockPing() {
	// 获得当前离明天早晨7点的时间距离, 即 每天早晨7点自动发送一条股市结果

	dingtalker := libs.NewDingtalker()
	go func() {
		s := libs.CrawlCarLimitJob()
		dingtalker.SendChatTextMessage(s)
	}()

	tn := time.Now().UTC().Add(8 * time.Hour)
	if tn.Weekday() == time.Monday || tn.Weekday() == time.Sunday {
		return
	}

	go func() {
		libs.CrawlStocksJob()
	}()
	time.Sleep(20000 * time.Millisecond)
	s, _ := libs.LastStock("baba")
	dingtalker.SendChatTextMessage(s.String())
}

func crawMovieJob() {
	go func() {
		libs.CrawlMovieJob()
		dingtalker := libs.NewDingtalker()
		dingtalker.SendChatLinkMessage("http://47.105.107.171/query?do=report", "http://47.105.107.171/sun.png", "3大视频网站热剧", "最新的优酷、腾讯、爱奇艺的热点电影近期上映及热度集中播报!")
	}()
}
