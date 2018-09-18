package libs

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func CrawlJob() {

	url := "http://m.iqiyi.com/vip/timeLine.html"

	t := time.Now()
	fo := fmt.Sprintf("%d-%02d-%02d-%02d", t.Year(), t.Month(), t.Day()-1, t.Hour()) //, t.Minute(), t.Second())
	fn := fmt.Sprintf("%d-%02d-%02d-%02d", t.Year(), t.Month(), t.Day(), t.Hour())   //, t.Minute(), t.Second())

	//fn := time.Now().Format("2006-01-02-15")
	fo = "./data/" + fo + "-iqiyi.html"
	err := os.Remove(fo) //删除24hours ago
	if err != nil {
		logs.Error(err)
		//输出错误详细信息
	}
	//	fn := time.Now().Format("2006-01-02-15")
	fn = "./data/" + fn + "-iqiyi.html"
	_, err = os.Stat(fn) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			logs.Error(err)
			return
		}
		logs.Debug("craw object:", fn)
		crawlByChrome(url, fn)
	}
}

func crawlByChrome(url string, fn string) {
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
	service, err := selenium.NewChromeDriverService("/usr/bin/chromedriver", 9515, opts...)
	if err != nil {
		logs.Error("Error starting the ChromeDriver server:", err)
	}

	defer service.Stop()

	// 调起chrome浏览器
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		logs.Error(err)
	}
	defer webDriver.Quit()

	// 这是目标网站留下的坑，不加这个在linux系统中会显示手机网页，每个网站的策略不一样，需要区别处理。
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
	pick_data := ""
	elem, err := webDriver.FindElement(selenium.ByClassName, "m-vip-timer-nav") //ByCSSSelector, "m-vip-timer-shaft")
	if err != nil {
		logs.Error(err)
	}
	output, err := elem.Text()
	if err != nil {
		logs.Error(err)
	}
	pick_data += output

	elem, err = webDriver.FindElement(selenium.ByClassName, "m-vip-timer-shaft")
	if err != nil {
		logs.Error(err)
	}

	melems, err := elem.FindElements(selenium.ByClassName, "border-left")
	if err != nil {
		logs.Error(err)
	}

	pick_data += "\n"
	for _, el := range melems {
		//	fmt.Println("\nNo.", k)
		pick_data += "\n"
		elem, err := el.FindElement(selenium.ByClassName, "title")
		if err != nil {
			// fmt.Println(err)
			pick_data += "\n时间：<等待排期>"
		} else {
			pick_data += "\n时间："
			s, _ := elem.Text()
			pick_data += s
		}
		elem, err = el.FindElement(selenium.ByClassName, "c-title")
		if err != nil {
			//fmt.Println(err)
		} else {
			pick_data += "\n影名："
			s, _ := elem.Text()
			pick_data += s
		}
		elem, err = el.FindElement(selenium.ByClassName, "album-history")
		if err != nil {
			//fmt.Println(err)
		} else {
			pick_data += "\n预约人数（万）："
			s, _ := elem.Text()
			s = strings.Replace(s, "万人已预约", "", -1)
			pick_data += s
		}
	}
	/*
		str, err := webDriver.PageSource()
		if err != nil {
			fmt.Println(err)
		}
	*/
	wc := []byte(pick_data)
	err = ioutil.WriteFile(fn, wc, 0644)
	if err != nil {
		logs.Error(err)
	}
}
