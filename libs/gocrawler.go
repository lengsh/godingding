package libs

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	//"io/ioutil"
	//	"os"
	"strconv"
	"strings"
	"time"
)

func CrawlMovieJob() {
	/*
		company := strings.ToUpper(com)
		switch company {
		case "IQIYI":
			crawlIqiyiByChrome()
		case "TX":
			crawlTxByChrome()
		}
	*/

	r := InitHuaweiCrawler()
	r.crawlIqiyiByChrome()
	r.crawlTxByChrome()
}

type Gocrawler struct {
	caps selenium.Capabilities
	opts []selenium.ServiceOption
}

func InitIOSCrawler() *Gocrawler {
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
			"--user-agent=Mozilla/5.0 (iPhone; CPU iPhone OS 11_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/11.0 Mobile/15E148 Safari/604.1"},
	}
	caps.AddChrome(chromeCaps)
	// 启动chromedriver，端口号可自定义

	// 调起chrome浏览器
	/*
		webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
		if err != nil {
			logs.Error(err)
	} */

	return &Gocrawler{caps, opts}
}

func InitHuaweiCrawler() *Gocrawler {
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
			"--user-agent=Mozilla/5.0 (Linux; Android 8.1.0; EML-AL00 Build/HUAWEIEML-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Mobile Safari/537.36"},
		// Mozilla/5.0 (iPhone; CPU iPhone OS 11_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/11.0 Mobile/15E148 Safari/604.1"},
	}
	caps.AddChrome(chromeCaps)
	// 启动chromedriver，端口号可自定义

	// 调起chrome浏览器
	/*
		webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
		if err != nil {
			logs.Error(err)
	} */

	return &Gocrawler{caps, opts}
}

func (r *Gocrawler) crawlIqiyiByChrome() {
	url := "http://m.iqiyi.com/vip/timeLine.html"

	service, err := selenium.NewChromeDriverService("/usr/bin/chromedriver", 9515, r.opts...)
	if err != nil {
		logs.Error("Error starting the ChromeDriver server:", err)
	}
	defer service.Stop()

	webDriver, err := selenium.NewRemote(r.caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		logs.Error(err)
	}
	defer webDriver.Quit()

	webDriver.AddCookie(&selenium.Cookie{
		Name:  "defaultJumpDomain",
		Value: "www",
	})

	// 导航到目标网站
	err = webDriver.Get(url)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to load page: %s\n", err))
	}
	//      fmt.Println(webDriver.Title())
	t := time.Now()
	fo := fmt.Sprintf("%d ", t.Year()) // , t.Month(), t.Day()-10, t.Hour()) //, t.Minute(), t.Second())
	elem, err := webDriver.FindElement(selenium.ByClassName, "m-vip-timer-shaft")
	if err != nil {
		logs.Error(err)
	}

	melems, err := elem.FindElements(selenium.ByClassName, "border-left")
	if err != nil {
		logs.Error(err)
	}

	for _, el := range melems {
		//	fmt.Println("\nNo.", k)
		var mo Movie
		mo.Company = "IQIYI"
		rt := ""
		elem, err := el.FindElement(selenium.ByClassName, "title")
		if err != nil {
			// fmt.Println(err)
			rt = "wait"
		} else {
			s, _ := elem.Text()
			rt = fo + strings.TrimSpace(s)
		}
		mo.Releasetime = rt

		elem, err = el.FindElement(selenium.ByClassName, "c-title")
		if err != nil {
			//fmt.Println(err)
		} else {
			s, _ := elem.Text()
			mo.Name = strings.TrimSpace(s)
		}

		elem, err = el.FindElement(selenium.ByClassName, "album-history")
		if err != nil {
			//fmt.Println(err)
		} else {
			s, _ := elem.Text()
			s = strings.Replace(s, "万人已预约", "", -1)
			value, err := strconv.ParseFloat(s, 32)
			if err != nil {
				// do something sensible
				value = 0
			}
			mo.Rate = float32(value)
		}
		fmt.Println(mo)
		//	mo.NewMovie()
	}
	/*
		str, err := webDriver.PageSource()
		if err != nil {
			fmt.Println(err)
		}
	*/
}

func (r *Gocrawler) crawlTxByChrome() {
	url := "http://film.qq.com/weixin/upcoming.html"

	service, err := selenium.NewChromeDriverService("/usr/bin/chromedriver", 9515, r.opts...)
	if err != nil {
		logs.Error("Error starting the ChromeDriver server:", err)
	}
	defer service.Stop()

	// 导航到目标网站

	webDriver, err := selenium.NewRemote(r.caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		logs.Error("Error to NewRmote webDriver ", err)
	}
	defer webDriver.Quit()

	webDriver.AddCookie(&selenium.Cookie{
		Name:  "defaultJumpDomain",
		Value: "www",
	})

	err = webDriver.Get(url)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to load page: %s\n", err))
	}
	//      fmt.Println(webDriver.Title())

	melems, err := webDriver.FindElements(selenium.ByClassName, "film_intro")
	if err != nil {
		logs.Error(err)
	}

	for _, el := range melems {

		var mo Movie
		mo.Company = "TX"

		elem, err := el.FindElement(selenium.ByClassName, "tit")
		if err != nil {
			fmt.Println(err)
		} else {
			s, _ := elem.Text()
			mo.Name = strings.TrimSpace(s)
		}

		elem, err = el.FindElement(selenium.ByClassName, "misc")
		if err != nil {
			fmt.Println(err)
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
			fmt.Println(err)
		} else {
			s, _ := elem.Text()
			val, err := strconv.ParseFloat(s, 32)
			if err == nil {
				mo.Rate = float32(val)
			}
		}
		fmt.Println(mo)
		//		mo.NewMovie()
	}
}
