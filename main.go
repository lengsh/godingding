package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/jasonlvhit/gocron"
	"github.com/lengsh/godingding/libs"
	"github.com/lengsh/godingding/utils"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var ( // main operation modes
	confFile = flag.String("config", "./config.json", "godingding configure file.")
)

func init() {

}

type Msg struct {
	Message string
	Scrumb  string
}

func usage() {
	// Fprintf allows us to print to a specifed file handle or stream
	fmt.Fprintf(os.Stderr, "Usage: %s [-flag xyz]\n", os.Args[0])
	flag.PrintDefaults()
}

func runInit() {
	flag.Usage = usage
	flag.Parse()
	// There is also a mandatory non-flag arguments
	if len(os.Args) < 2 {
		usage()
	}

	utils.ConfigInit(*confFile)
	libs.DBinit(utils.DBConfig.URL, utils.DBConfig.MaxIdleConns, utils.DBConfig.MaxOpenConns)

	sLogDaily := "false"
	if utils.ServerConfig.LogDaily {
		sLogDaily = "true"
	}
	sl := fmt.Sprintf("{\"filename\":\"%s\",\"maxlines\":%d,\"maxsize\":%d,\"daily\":%s,\"maxdays\":%d}", utils.ServerConfig.LogFile, utils.ServerConfig.LogMaxLines, utils.ServerConfig.LogMaxSize, sLogDaily, utils.ServerConfig.LogMaxDays)
	// logs.SetLogger(logs.AdapterFile, `{"filename":"./log/godingding.log","maxlines":10000,"maxsize":102400,"daily":true,"maxdays":2    }`)
	logs.SetLogger(logs.AdapterFile, sl)
	logs.EnableFuncCallDepth(utils.ServerConfig.LogEnableDepth) //   true)
	logs.SetLogFuncCallDepth(utils.ServerConfig.LogDepth)       // 3 )
	switch utils.ServerConfig.LogLevel {
	case "DEBUG":
		logs.SetLevel(logs.LevelDebug)
	case "ERROR":
		logs.SetLevel(logs.LevelError)

	case "WARNING":
		logs.SetLevel(logs.LevelWarning)
	default:
		logs.SetLevel(logs.LevelError)

	}
}

func main() {
	runInit()
	http.HandleFunc("/", firstPage)    //设置访问的路由
	http.HandleFunc("/send", send)     //设置访问的路由
	http.HandleFunc("/query", query)   //设置访问的路由
	http.HandleFunc("/stock", stock)   //设置访问的路由
	http.HandleFunc("/help", help)     //设置访问的路由
	http.HandleFunc("/jsonp", jsonApi) //设置访问的路由

	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("/usr/share/nginx/html/doc"))))

	scheduler := gocron.NewScheduler()

	job0 := scheduler.Every(1).Day().At("07:01")
	job0.Do(carLimit)

	job1 := scheduler.Every(1).Day().At("07:09")
	job1.Do(stockPing)

	job2 := scheduler.Every(1).Day().At("09:09")
	job2.Do(stockPing)

	job3 := scheduler.Every(1).Day().At("10:49")
	job3.Do(crawMovieJob)

	job4 := scheduler.Every(1).Day().At("20:01")
	job4.Do(crawMovieJob)
	//scheduler.Every(1).Day().Do(crawJob)
	scheduler.Every(1).Hour().Do(crawResouJob)
	scheduler.Start()

	port := fmt.Sprintf(":%d", utils.ServerConfig.Port)
	fmt.Println("http server listen to port:", port)

	s := &http.Server{
		Addr:         port,
		Handler:      http.DefaultServeMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		//      MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			//			log.Println(err)
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	//    signal.Notify(ch, os.Interrupt)

	// Block until we receive our signal.
	// Handle SIGINT and SIGTERM.
	<-ch
	//  log.Println(<-ch)
	// Stop the service gracefully.
	// log.Println(s.Shutdown(nil))
	// Wait gorotine print shutdown message

	wait := time.Second * 5
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	s.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("graceful shutdown -->done!")
	os.Exit(0)

	/*
		err := http.ListenAndServe(":8081", nil) //设置监听的端口
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		} */
}

func queryStock(w http.ResponseWriter, r *http.Request) {
	qs := libs.QueryStock(10)
	t, _ := template.ParseFiles("view/stock.gtpl", "view/header.gtpl", "view/footer.gtpl")
	err := t.Execute(w, qs)
	if err != nil {
		logs.Error(err.Error())
	}
}

func resouReport(w http.ResponseWriter, r *http.Request) {
	var qy []libs.TouTiao
	key := time.Now().Format("2006-01-02")

	if value, ret := libs.GetKVStore("RESOU", key); ret {
		err := json.Unmarshal([]byte(value), &qy)
		if err != nil {
			logs.Error("Unmarshal failed, ", err)
			qy = nil
		}
	} else {
		go func() {
			qy := libs.GrabToutiaoProcess()
			data, err := json.Marshal(qy)
			if err == nil {
				libs.SetKVStore("RESOU", key, string(data))
			}
		}()
	}
	words := "<h2>" + libs.FetchKeyWords() + "</h2>"
	t, _ := template.New("resou.gtpl").Funcs(template.FuncMap{
		"ResouKeyWordsA": func() template.HTML {
			return template.HTML(words)
		},
		"ResouKeyWordsB": func() template.HTML {
			return template.HTML("<h4>HelloB </h4>")
		},
	}).ParseFiles("view/resou.gtpl", "view/header.gtpl", "view/footer.gtpl")

	//	t, _ := template.ParseFiles("view/resou.gtpl", "view/header.gtpl", "view/footer.gtpl")
	err := t.Execute(w, qy)
	if err != nil {
		logs.Debug(err.Error())
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
		fmt.Println(err)
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
				scrum_new := utils.CreateScrumb("send")
				if scrum_new == n[0] {
					dingtalker := utils.NewDingtalker()
					sm := m[0] // template.HTMLEscapeString(m[0])
					dingtalker.SendChatTextMessage(sm)
				}
			}
		}
	}
	scr := utils.CreateScrumb("send")
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
			} else if m[0] == "resou" {
				key = 4
			}
		}
	}
	switch key {
	case 4:
		resouReport(w, r)
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
				fmt.Println(ss)
				dingtalker := utils.NewDingtalker()
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
					s = "'xianxing' or 'car limit' -- 汽车限行信息\n'update stock' -- 股票系统数据更新\n'update movie'  -- 影视信息数据更新\n'BABA'...etc  -- 特定股票单独更新与咨询指令"
				}
				ss := fmt.Sprintf("@%s\n%s", senderNick, s)
				dingtalker := utils.NewDingtalker()
				dingtalker.SendRobotTextMessage(ss)
			}
		}()
	}

	if r.Method == "GET" {
		// return data/xxx.data
	}
}

func jsonApi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		if n, ok2 := r.Form[".scrumb"]; ok2 {
			scrum_new := utils.CreateScrumb("jsonapi")
			if scrum_new == n[0] {
				rets := ""
				if f, ok2 := r.Form["func"]; ok2 {
					p := "BABA"
					if p0, ok2 := r.Form["para"]; ok2 {
						p = p0[0]
					}
					fc := strings.ToUpper(f[0])
					switch fc {
					case "STOCK":
						rets = libs.StockLineJson(p)
					case "STOCKSUM":
						rets = libs.StockMarketCapJson("2017-09-01")
					}
				}
				fmt.Fprintf(w, rets)
			} else {
				fmt.Fprintf(w, "Error invoker, Nothing to do!")
			}
		} else {
			fmt.Fprintf(w, "Error invoke, no privilate to read data!")
		}
	} else {
		fmt.Fprintf(w, "Error invoke, no GET parameter !")
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

func carLimit() {
	// 获得当前离明天早晨7点的时间距离, 即 每天早晨7点自动发送一条股市结果

	dingtalker := utils.NewDingtalker()
	s := libs.CrawlCarLimitJob()
	dingtalker.SendChatTextMessage(s)
}

func stockPing() {
	// 获得当前离明天早晨7点的时间距离, 即 每天早晨7点自动发送一条股市结果

	tn := time.Now().UTC().Add(8 * time.Hour)
	if tn.Weekday() == time.Monday || tn.Weekday() == time.Sunday {
		return
	}
	go func() {
		libs.CrawlStocksJob()
		i := rand.Intn(100)
		if i%2 == 0 {
			o, _ := libs.LastStock("baba")
			dingtalker := utils.NewDingtalker()
			dingtalker.SendChatTextMessage(o.String())
		}
	}()
}

func crawMovieJob() {
	go func() {
		libs.CrawlMovieJob()
		dingtalker := utils.NewDingtalker()
		dingtalker.SendChatLinkMessage("http://47.105.107.171/query?do=report", "http://47.105.107.171/sun.png", "3大视频网站热剧", "最新的优酷、腾讯、爱奇艺的热点电影近期上映及热度集中播报!")

		tn := time.Now().UTC().Add(8 * time.Hour)
		if tn.Weekday() == time.Monday || tn.Weekday() == time.Sunday {
			return
		}
		libs.CrawlStocksJob()
	}()
}

func crawResouJob() {
	go func() {
		key := time.Now().Format("2006-01-02")
		tt := libs.GrabToutiaoProcess()
		libs.PickKeyWords(tt)
		data, err := json.Marshal(tt)
		if err == nil {
			libs.SetKVStore("RESOU", key, string(data))
		} else {
			logs.Error("Error to Grab Resou !!!!!!!!!")
		}
	}()
}
