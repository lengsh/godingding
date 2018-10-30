package libs

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	//	"io/ioutil"
	//	"os"
	"strconv"
	"strings"
	"time"
)

func CrawlStocksJob() {
	r := NewCrawler()
	defer r.ReleaseCrawler()
	stocks := "BABA,FB,MSFT,AMZN,AAPL,TSLA,BIDU,NVDA,GOOGL,WB"
	stocksv := strings.Split(stocks, ",")

	for _, v := range stocksv {
		logs.Debug(v)
		r.crawlStockByChrome(v)
		time.Sleep(10000 * time.Millisecond)
	}

	///////////////////////////////////////////////////
	i := 0
	for i < 10 {
		nvect := make(map[string]string)
		sv := QueryTodayStock()
		if len(sv) != len(stocksv) {
			for _, ns := range stocksv {
				b := false
				for _, op := range sv {
					if ns == op.Name {
						b = true
					}
				}
				if !b {
					nvect[ns] = ns
				}
			}
			if len(nvect) == 0 { // never to be run !!!!!
				break
			}
			logs.Debug("the ", i, " times to grab.....")
			for _, v := range nvect {
				logs.Debug(v)
				switch i % 3 {
				case 0:
					r.crawlStockBy163Chrome(v)
				case 1:
					r.crawlStockByBaiduChrome(v)
				case 2:
					r.crawlStockByChrome(v)
				}
				time.Sleep(100000 * time.Millisecond)
			}
		} else {
			logs.Debug("Yes,data is ready, break!")
			break
		}
		time.Sleep(10000000 * time.Millisecond)
		i++
	}
}

func CrawlStockJob(sk string) string {
	r := NewCrawler()
	defer r.ReleaseCrawler()

	s := r.crawlStockByChrome(sk)
	if strings.Contains(s, "error") {
		return r.crawlStockByBaiduChrome(sk)
	}
	return s
}

func CrawlCarLimitJob() string {
	r := NewCrawler()
	defer r.ReleaseCrawler()

	s := r.crawlBaiduXXByChrome()
	if strings.Contains(s, "error") {
		return r.crawlSogouXXByChrome()
	}
	return s
}

func CrawlMovieJob() {
	r := NewCrawler()
	defer r.ReleaseCrawler()
	r.crawlIqiyiByChrome()
	r.crawlTxByChrome()
	r.crawlYoukuByChrome()

	NewMylog("sys", "Update")

	mvs := QueryTopMovies("IQIYI", 20)
	for _, mv := range mvs {
		fr := r.crawlDoubanByChrome(mv.Name)
		if fr != mv.Douban {
			mv.Douban = fr
			UpdateMovie(mv)
		}
		time.Sleep(2000 * time.Millisecond)
	}

	mvs = QueryTopMovies("TX", 20)
	for _, mv := range mvs {
		fr := r.crawlDoubanByChrome(mv.Name)
		if fr != mv.Douban {
			mv.Douban = fr
			UpdateMovie(mv)
		}
		time.Sleep(2000 * time.Millisecond)
	}

	mvs = QueryTopMovies("YOUKU", 20)
	for _, mv := range mvs {
		fr := r.crawlDoubanByChrome(mv.Name)
		if fr != mv.Douban {
			mv.Douban = fr
			UpdateMovie(mv)
		}
		time.Sleep(2000 * time.Millisecond)
	}
}

func UpdateDouban() {
	r := NewCrawler()
	defer r.ReleaseCrawler()

	mvs := QueryZeroDouban(100)
	for _, mv := range mvs {
		logs.Debug("craw :", mv.Name, "; length=", len(mv.Name), " ?=", len(strings.TrimSpace(mv.Name)))
		fr := r.crawlDoubanByChrome(strings.TrimSpace(mv.Name))
		if fr != mv.Douban {
			mv.Douban = fr
			UpdateMovie(mv)
		}
		time.Sleep(10000 * time.Millisecond)
	}
}

type GoCrawler struct {
	service   *selenium.Service
	webDriver selenium.WebDriver
}

func NewCrawler() *GoCrawler {
	// StartChrome 启动谷歌浏览器headless模式
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName": "chrome",
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
			//      "--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
			//	"--user-agent=Mozilla/5.0 (Linux; Android 8.1.0; EML-AL00 Build/HUAWEIEML-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Mobile Safari/537.36"},
			"--user-agent=Mozilla/5.0 (iPhone; CPU iPhone OS 12_0_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/11.0 Mobile/15E148 Safari/604.1"},
	}
	caps.AddChrome(chromeCaps)
	// 启动chromedriver，端口号可自定义

	// 调起chrome浏览器
	service, err := selenium.NewChromeDriverService("/usr/bin/chromedriver", 9515, opts...)
	if err != nil {
		logs.Error("Error starting the ChromeDriver server:", err)
	}

	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		logs.Error(err)
	}

	return &GoCrawler{service, webDriver}
}

func (r *GoCrawler) ReleaseCrawler() {
	defer r.service.Stop()
	defer r.webDriver.Quit()
}

func (r *GoCrawler) crawlIqiyiByChrome() {
	url := "http://m.iqiyi.com/vip/timeLine.html"

	r.webDriver.AddCookie(&selenium.Cookie{
		Name:  "defaultJumpDomain",
		Value: "www",
	})

	// 导航到目标网站
	err := r.webDriver.Get(url)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to load page: %s\n", err))
		es := "[WARNING] " + url + " May be shutdown, please make true now!"
		dingtalker := NewDingtalker()
		dingtalker.SendRobotTextMessage(es)
		return
	}
	//      fmt.Println(webDriver.Title())
	t := time.Now()
	fo := fmt.Sprintf("%d ", t.Year()) // , t.Month(), t.Day()-10, t.Hour()) //, t.Minute(), t.Second())
	elem, err := r.webDriver.FindElement(selenium.ByClassName, "m-vip-timer-shaft")
	if err != nil {
		logs.Error(err)
		return
	}

	melems, err := elem.FindElements(selenium.ByClassName, "border-left")
	if err != nil {
		logs.Error(err)
		return
	}

	for _, el := range melems {
		//	fmt.Println("\nNo.", k)
		var mo Movie
		mo.Company = "IQIYI"
		rt := ""
		elem, err := el.FindElement(selenium.ByClassName, "title")
		if err != nil {
			logs.Debug(err)
			rt = "wait"
		} else {
			s, _ := elem.Text()
			rt = fo + strings.TrimSpace(s)
		}
		mo.Releasetime = rt

		elem, err = el.FindElement(selenium.ByClassName, "c-title")
		if err != nil {
			logs.Debug(err)
		} else {
			s, _ := elem.Text()
			mo.Name = strings.TrimSpace(s)
		}

		elem, err = el.FindElement(selenium.ByClassName, "album-history")
		if err != nil {
			logs.Debug(err)
		} else {
			s, _ := elem.Text()
			s = strings.Replace(s, "万人已预约", "", -1)
			value, err := strconv.ParseFloat(s, 32)
			if err != nil {
				logs.Debug(err)
				value = 0
			}
			mo.Rate = float32(value)
		}
		//	fmt.Println(mo)
		mo.NewMovie()
	}
	/*
		str, err := webDriver.PageSource()
		if err != nil {
			fmt.Println(err)
		}
	*/
}

func (r *GoCrawler) crawlTxByChrome() {
	url := "http://film.qq.com/weixin/upcoming.html"

	r.webDriver.AddCookie(&selenium.Cookie{
		Name:  "defaultJumpDomain",
		Value: "www",
	})

	err := r.webDriver.Get(url)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to load page: %s\n", err))
		es := "[WARNING] " + url + " May be shutdown, please make true now!"
		dingtalker := NewDingtalker()
		dingtalker.SendRobotTextMessage(es)
		return
	}
	//      fmt.Println(webDriver.Title())

	melems, err := r.webDriver.FindElements(selenium.ByClassName, "film_intro")
	if err != nil {
		logs.Error(err)
	}

	for _, el := range melems {

		var mo Movie
		mo.Company = "TX"

		elem, err := el.FindElement(selenium.ByClassName, "tit")
		if err != nil {
			logs.Debug(err)
		} else {
			s, _ := elem.Text()
			mo.Name = strings.TrimSpace(s)
		}

		elem, err = el.FindElement(selenium.ByClassName, "misc")
		if err != nil {
			logs.Debug(err)
		} else {
			s, _ := elem.Text()
			ts := strings.Split(s, " ")
			mo.Releasetime = "wait"
			if len(ts) >= 2 {
				t, _ := time.Parse("2006-01-02", ts[0])
				mo.Releasetime = fmt.Sprintf("%d %02d月%02d日", t.Year(), t.Month(), t.Day())
			}
		}

		elem, err = el.FindElement(selenium.ByClassName, "score_wrap")
		if err != nil {
			logs.Debug(err)
		} else {
			s, _ := elem.Text()
			val, err := strconv.ParseFloat(s, 32)
			if err == nil {
				mo.Rate = float32(val)
			}
		}
		//	fmt.Println(mo)
		mo.NewMovie()
	}
}

func (r *GoCrawler) crawlYoukuByChrome() {

	//url := "https://vip.youku.com/vips/index.html"
	url := "https://h5.vip.youku.com"
	r.webDriver.AddCookie(&selenium.Cookie{
		Name:  "defaultJumpDomain",
		Value: "www",
	})

	err := r.webDriver.Get(url)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to load page: %s\n", err))
		es := "[WARNING] " + url + " May be shutdown, please make true now!"
		dingtalker := NewDingtalker()
		dingtalker.SendRobotTextMessage(es)
		return
	}
	//  hot vedio
	melem, err := r.webDriver.FindElement(selenium.ByClassName, "movie-lists")
	if err != nil {
		logs.Error(err)
	} else {
		melems, err := melem.FindElements(selenium.ByTagName, "li")
		if err != nil {
			logs.Error(err)
		} else {
			//	fmt.Println("近期热播,count = ", len(melems))
			tn := time.Now().UTC().Add(8 * time.Hour)
			for _, mel := range melems {
				slin, err := mel.Text()
				if err != nil {
					logs.Error(err)
				} else {
					v := strings.Split(slin, "\n")
					if len(v) == 4 {
						var mo Movie
						mo.Company = "YOUKU"
						val, _ := strconv.ParseFloat(v[1], 32)
						mo.Rate = float32(val)
						mo.Name = strings.TrimSpace(v[2])
						mo.Releasetime = fmt.Sprintf("%d %02d月%02d日", tn.Year(), tn.Month(), tn.Day())
						//mo.Releasetime = "running"
						mo.NewMovie()
					} else {
						logs.Error("不合规数据:", v)
					}
				}
			}
		}
	}
	//  comming soon
	melem, err = r.webDriver.FindElement(selenium.ByClassName, "movielist-container") //swiper-container-book")
	if err != nil {
		logs.Error(err)
		return
	}

	//	fmt.Println(melem.Text())

	melems, err := melem.FindElements(selenium.ByTagName, "dl")
	if err != nil {
		logs.Error(err)
		return
	}

	for _, el := range melems {
		var mo Movie
		mo.Company = "YOUKU"

		elem, err := el.FindElement(selenium.ByTagName, "dt")
		if err != nil {
			logs.Debug(err)
		} else {

			el1, err := elem.FindElement(selenium.ByClassName, "score")
			if err != nil {
				logs.Debug(err)
				mo.Rate = 0
			} else {
				s, _ := el1.Text()
				val, _ := strconv.ParseFloat(s, 32)
				mo.Rate = float32(val)
			}
		}

		elem, err = el.FindElement(selenium.ByTagName, "dd")
		if err != nil {
			logs.Debug(err)
		} else {

			el1, err := elem.FindElement(selenium.ByTagName, "h3")
			if err != nil {
				logs.Debug(err)
			} else {
				mo.Name, _ = el1.Text()
			}
			el2, err := elem.FindElement(selenium.ByTagName, "p")
			if err != nil {
				logs.Debug(err)
			} else {
				s, _ := el2.Text()
				s = strings.TrimSpace(s)
				t, _ := time.Parse("2006-01-02", s)
				if t.Year() >= time.Now().Year() {
					mo.Releasetime = fmt.Sprintf("%d %02d月%02d日", t.Year(), t.Month(), t.Day())
				} else {
					mo.Releasetime = "wait"
				}
			}
		}
		//	fmt.Println(mo)
		mo.NewMovie()
	}
}

func (r *GoCrawler) crawlDoubanByChrome(mv string) float32 {
	mv = strings.TrimSpace(mv)
	url := fmt.Sprintf("https://www.douban.com/search?cat=1002&q=%s", mv)
	/*
		r.webDriver.AddCookie(&selenium.Cookie{
			Name:  "defaultJumpDomain",
			Value: "www",
		})
	*/
	err := r.webDriver.Get(url)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to load page: %s\n", err))
		es := "[WARNING] " + url + " May be shutdown, please make true now!"
		dingtalker := NewDingtalker()
		dingtalker.SendRobotTextMessage(es)
		return 0
	}
	//      fmt.Println(webDriver.Title())

	melem, err := r.webDriver.FindElement(selenium.ByClassName, "search-result")
	if err != nil {
		logs.Error(err)
		return 0
	}

	melems, err := melem.FindElements(selenium.ByClassName, "content")
	if err != nil {
		logs.Error(err)
		return 0
	}

	var fr float32 = 0
	for _, el := range melems {
		logs.Debug(el.Text())
		sn := ""
		elem, err := el.FindElement(selenium.ByClassName, "title")
		if err != nil {
			logs.Debug(err)
		} else {
			s, _ := elem.Text()
			s = strings.Replace(s, "[电影]", "", 1)
			s = strings.TrimSpace(s)
			sv := strings.Split(s, "\n")
			if len(sv) > 0 {
				sn = strings.TrimSpace(sv[0])
			}
			logs.Debug(s, " >>> name= ", sn)
		}

		elem, err = el.FindElement(selenium.ByClassName, "rating-info")
		if err != nil {
			logs.Debug(err)
		} else { //rating_nums

			s, _ := elem.Text()
			sv := strings.Split(s, " ")
			logs.Debug(s, ">>> rate =", sv)
			if len(sv) > 1 {

				f, _ := strconv.ParseFloat(sv[0], 32)
				fr = float32(f)

				/*
					ns := strings.Replace(sv[1], "(", "", -1)
					ns = strings.Replace(ns, "人评价)", "", -1)
					fmt.Println(ns)
				*/
			} else {
			}
		}
		logs.Debug("mv=", mv, " >>> sn=", sn, "rate = ", fr)
		if mv == sn {
			return fr
		}
	}
	return fr
}

//////////
func (r *GoCrawler) crawlStockBy163Chrome(sID string) string {
	logs.Debug("try to crawl ", sID)
	mv := strings.ToUpper(strings.TrimSpace(sID))
	url := fmt.Sprintf("http://quotes.money.163.com/usstock/%s.html", mv)

	r.webDriver.AddCookie(&selenium.Cookie{
		Name:  "a",
		Value: "www",
	})
	err := r.webDriver.Get(url)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to load page: %s\n", err))
		es := "[!!!] " + url + " May be shutdown, please make true now!"
		logs.Error(es)
		return "0 error"
	}

	melem, err := r.webDriver.FindElement(selenium.ByClassName, "banner")
	if err != nil {
		logs.Error(err)
		return "1 error"
	}
	scur, err := melem.Text()
	if err != nil {
		logs.Error(err)
		return "2 error"
	}
	vs := strings.Split(scur, "\n")
	if len(vs) < 10 {
		logs.Error("数据页面结构已经修改，需要重新编码！！！！！！")
		return "3 error"
	}

	var mo Stockorm
	mo.Name = mv

	ns := strings.Split(vs[2], " ")
	if len(ns) < 2 {
		return "4 error"
	}

	f, _ := strconv.ParseFloat(ns[1], 32)
	mo.EndPrice = float32(f)

	ns = strings.Split(vs[6], " ")
	if len(ns) < 2 {
		return "4 error"
	}

	s := ns[0]
	if len(s) > 6 {
		f, _ := strconv.ParseFloat(s[6:], 32)
		mo.StartPrice = float32(f)
	}
	s = ns[1]
	if len(s) > 7 {
		f, _ := strconv.ParseFloat(s[6:], 32)
		mo.HighPrice = float32(f)
	}
	ns = strings.Split(vs[7], " ")
	if len(ns) < 5 {
		return "4 error"
	}
	s = ns[1]
	if len(s) > 7 {
		f, _ = strconv.ParseFloat(s[6:], 32)
		mo.LowPrice = float32(f)
	}
	// ????
	mo.TradeFounds = 0.0
	mo.TradeStock = 0.0

	s = ns[4]
	bb := strings.Contains(s, "万亿")
	if bb {
		if len(s)-8 > 9 { //  len("(万亿)") == 8
			f, _ = strconv.ParseFloat(s[9:len(s)-5], 32)
			mo.MarketCap = float32(f) * 10000
		}
	} else {
		if len(s)-5 > 9 { //  len("(亿)") == 5
			f, _ = strconv.ParseFloat(s[9:len(s)-5], 32)
			mo.MarketCap = float32(f)
		}
	}

	mo.CreateDate = time.Now().UTC().Add(8 * time.Hour)
	if mo.MarketCap > 0 {
		mo.NewStock()
	} else {
		logs.Error("MarketCap is Zero !!!!!!!!!!!!!! ")
	}
	return mo.String()
}

//////////
func (r *GoCrawler) crawlStockByBaiduChrome(sID string) string {
	logs.Debug("try to crawl ", sID)
	mv := strings.ToUpper(strings.TrimSpace(sID))
	url := fmt.Sprintf("https://www.baidu.com/s?wd=%s&rsv_spt=1", mv)
	r.webDriver.AddCookie(&selenium.Cookie{
		Name:  "a",
		Value: "www",
	})
	err := r.webDriver.Get(url)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to load page: %s\n", err))
		es := "[!!!] " + url + " May be shutdown, please make true now!"
		logs.Error(es)
		return "0 error"
	}

	melem, err := r.webDriver.FindElement(selenium.ByClassName, "result")
	if err != nil {
		logs.Error(err)
		return "error"
	}
	scur, err := melem.Text()
	if err != nil {
		logs.Error(err)
		return "error"
	}
	vs := strings.Split(scur, "\n")
	if len(vs) < 31 {
		logs.Error("数据页面结构已经修改，需要重新编码！！！！！！")
		return "3 error"
	}

	var mo Stockorm
	mo.Name = mv

	f, _ := strconv.ParseFloat(vs[21], 32)
	mo.HighPrice = float32(f)

	f, _ = strconv.ParseFloat(vs[23], 32)
	mo.LowPrice = float32(f)
	f, _ = strconv.ParseFloat(vs[17], 32)
	mo.StartPrice = float32(f)

	nvs := strings.Split(vs[1], "+")
	f, _ = strconv.ParseFloat(nvs[0], 32)
	mo.EndPrice = float32(f)
	// ????
	mo.TradeFounds = 0.0

	s := strings.Replace(vs[25], "万", "", -1)
	s = strings.Replace(s, "亿", "", -1)
	f, _ = strconv.ParseFloat(s, 32)
	mo.TradeStock = float32(f)

	bb := strings.Contains(vs[31], "万亿")
	s = strings.Replace(vs[31], "万", "", -1)
	s = strings.Replace(s, "亿", "", -1)
	f, _ = strconv.ParseFloat(s, 32)
	mo.MarketCap = float32(f)
	if bb {
		mo.MarketCap = mo.MarketCap * 10000
	}

	mo.CreateDate = time.Now().UTC().Add(8 * time.Hour)
	if mo.MarketCap > 0 {
		mo.NewStock()
	} else {
		logs.Error("MarketCap is Zerooooooooooooooooooo !!")
	}
	return mo.String()
}

////////
func (r *GoCrawler) crawlStockByChrome(sID string) string {
	logs.Debug("try to crawl ", sID)
	mv := strings.ToUpper(strings.TrimSpace(sID))
	url := fmt.Sprintf("https://www.futunn.com/quote/stock?m=us&code=%s", mv)
	r.webDriver.AddCookie(&selenium.Cookie{
		Name:  "a",
		Value: "www",
	})
	err := r.webDriver.Get(url)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to load page: %s\n", err))
		es := "[!!!] " + url + " May be shutdown, please make true now!"
		// dingtalker := NewDingtalker()
		// dingtalker.SendRobotTextMessage(es)
		logs.Error(es)
		return "0 error"
	}

	melem, err := r.webDriver.FindElement(selenium.ByClassName, "price01")
	if err != nil {
		logs.Error(err)
		return "1 error"
	}
	scur, err := melem.Text()
	if err != nil {
		logs.Error(err)
		return "2 error"
	}

	melem, err = r.webDriver.FindElement(selenium.ByClassName, "listBar")
	if err != nil {
		logs.Error(err)
		fmt.Println("error?")
		return "3 error"
	}

	var mo Stockorm
	mo.Name = mv
	sbuf, err2 := melem.Text()
	if err2 == nil {
		sv := strings.Split(sbuf, "\n")
		if len(sv) == 10 {
			s := sv[0]
			sbv := strings.Split(s, "：")
			if len(sbv) != 2 {
				goto RETURN
			}
			f, _ := strconv.ParseFloat(sbv[1], 32)
			mo.HighPrice = float32(f)

			s = sv[1]
			sbv = strings.Split(s, "：")
			if len(sbv) != 2 {
				goto RETURN
			}
			f, _ = strconv.ParseFloat(sbv[1], 32)
			mo.LowPrice = float32(f)

			s = sv[2]
			sbv = strings.Split(s, "：")
			if len(sbv) != 2 {
				goto RETURN
			}
			f, _ = strconv.ParseFloat(sbv[1], 32)
			mo.StartPrice = float32(f)
			/*
				s = sv[3]
				sbv = strings.Split(s, "：")
				if len(sbv) != 2 {
					goto RETURN
				}
				f, _ = strconv.ParseFloat(sbv[1], 32)
				mo.EndPrice = float32(f)
			*/
			f, _ = strconv.ParseFloat(scur, 32)
			mo.EndPrice = float32(f)

			s = sv[4]
			sbv = strings.Split(s, "：")
			if len(sbv) != 2 {
				goto RETURN
			}
			s = strings.Replace(sbv[1], "万", "", -1)
			s = strings.Replace(s, "亿", "", -1)
			f, _ = strconv.ParseFloat(s, 32)
			mo.TradeFounds = float32(f)

			s = sv[5]
			sbv = strings.Split(s, "：")
			if len(sbv) != 2 {
				goto RETURN
			}
			s = strings.Replace(sbv[1], "万", "", -1)
			s = strings.Replace(s, "亿", "", -1)
			f, _ = strconv.ParseFloat(s, 32)
			mo.TradeStock = float32(f)

			s = sv[7]
			sbv = strings.Split(s, "：")
			if len(sbv) != 2 {
				goto RETURN
			}

			bb := strings.Contains(s, "万亿")

			s = strings.Replace(sbv[1], "万", "", -1)
			s = strings.Replace(s, "亿", "", -1)
			f, _ = strconv.ParseFloat(s, 32)
			mo.MarketCap = float32(f)
			if bb {
				mo.MarketCap = mo.MarketCap * 10000
			}
			mo.CreateDate = time.Now().UTC().Add(8 * time.Hour)
			/////////
			///////////
			if mo.MarketCap > 0 {
				mo.NewStock()
			} else {
				logs.Error("MarketCap is zeroOOOOOOOOOOOOOOooooooooooooo !!")
			}
			return mo.String()
			//	fmt.Println(mo)

		} else {
			logs.Error("Error to parser")
			fmt.Println("error ??")
		}
	} else {
		logs.Error(err2, " Error to parser")
		fmt.Println("error ??")
	}
RETURN:
	logs.Error("something wrong with ", sID)
	return "error"
}

//////////////
func (r *GoCrawler) crawlSogouXXByChrome() string {
	url := "https://m.sogou.com/web/searchList.jsp?keyword=限行尾号&wm=3206"
	err := r.webDriver.Get(url)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to load page: %s\n", err))
		// es := "[WARNING] " + url + " May be shutdown, please make true now!"
		return "0 error"
	}
	melem, err := r.webDriver.FindElement(selenium.ByClassName, "vr-limit180417")
	if err != nil {
		logs.Error(err)
		return "1 error"
	}

	sn, _ := melem.Text()
	v := strings.Split(sn, "\n")
	/*
		       	       for k,vv := range v {
				       		       fmt.Println(k,"=>",vv)
		       	       }*/
	rets := ""
	if len(v) >= 8 {
		rets = fmt.Sprintf("%s\n今天(%s)限行：%s\n明天(%s)限行：%s", v[6], v[1], v[2], v[4], v[5])
	} else {
		rets = "2 error!"
	}
	return rets
}

///////////////
func (r *GoCrawler) crawlBaiduXXByChrome() string {
	url := "https://www.baidu.com/from=844b/s?word=%E5%8C%97%E4%BA%AC%E9%99%90%E5%8F%B7&sa=tb&ms=1"
	err := r.webDriver.Get(url)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to load page: %s\n", err))
		// es := "[WARNING] " + url + " May be shutdown, please make true now!"
		return "0 error"
	}

	melem, err := r.webDriver.FindElement(selenium.ByClassName, "s-cluster-container")
	if err != nil {
		logs.Error(err)
		return "1 error"
	}

	sn, _ := melem.Text()
	v := strings.Split(sn, "\n")
	rets := ""
	if len(v) >= 8 {
		rets = fmt.Sprintf("%s\n%s%s%s%s%s\n%s%s%s%s%s", v[7], v[1], "(", v[2], "):", v[3], v[4], "(", v[5], "):", v[6])
	} else {
		return "2 error"
	}
	return rets
}
