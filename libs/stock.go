package libs

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/benbjohnson/phantomjs"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Crawler_163(st string) string {
	sm := time.Now().Format("2006-01-02 15:04:05")
	if err := phantomjs.DefaultProcess.Open(); err != nil {
		logs.Error(err)
		return sm
	}
	defer phantomjs.DefaultProcess.Close()
	page, err := phantomjs.CreateWebPage()
	if err != nil {
		logs.Error(err)
		return sm
	}
	stk := strings.ToUpper(st)
	surl := fmt.Sprintf("http://quotes.money.163.com/usstock/%s.html#US1a01", stk)

	if err := page.Open(surl); err != nil {
		logs.Error(err)
		return sm
	}

	if content, err := page.Content(); err == nil {
		idx1 := strings.Index(content, "<div class=\"stock_info\">")
		idx2 := strings.Index(content, "<div class=\"stock_nav_bar\">")
		s1 := content[idx1:idx2]

		/*
			// <div class="time">
			idx1 = strings.Index(s1, "<div class=\"time\">")
			idx2 = strings.Index(s1, "<div class=\"stock_detail\">")
			s2 := s1[idx1+30:idx2]
			idx1 = strings.Index(s2, "\">")
			idx2 = strings.Index(s2, "</span>")
			sTime := s2[idx1+2:idx2]
		*/
		idx1 = strings.Index(s1, "<div class=\"time\">")
		idx2 = strings.Index(s1, "<div class=\"stock_detail\">")
		s2 := s1[idx1+30 : idx2]
		idx1 = strings.Index(s2, "\">")
		idx2 = strings.Index(s2, "</span>")
		sTime := s2[idx1+2 : idx2]
		if len(sTime) > 4 {
			logs.Debug("163'Datetime = ", sTime)
			return sTime
		} else {
			return sm
		}

	}
	return sm
}

func Crawler_Futu(st string, dt string) Stock {
	if len(dt) < 10 {
		dt = time.Now().Format("2006-01-02 15:04:05")
		logs.Error("datetime Error ,reset current datetime = ", dt)
	}
	stk := strings.ToUpper(st)
	surl := fmt.Sprintf("https://www.futunn.com/quote/stock?m=us&code=%s", stk)

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(surl)
	if err != nil {
		log.Println(err)
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}

	str := string(body[:])

	sStart := "<span class=\"price01\">"
	sEnd := "<script type=\"text/template\" id=\"basicQuoteTpl\">"
	idx1 := strings.Index(str, sStart)
	idx2 := strings.Index(str, sEnd)
	s1 := str[idx1+len(sStart) : idx2]
	sEnd = "</span>"
	idx2 = strings.Index(s1, sEnd)
	sCur := s1[0:idx2]

	sStart = "最　高："
	sEnd = "<p>最　低："
	idx1 = strings.Index(s1, sStart)
	idx2 = strings.Index(s1, sEnd)
	s2 := s1[idx1+len(sStart) : idx2]
	idx1 = strings.Index(s2, "\">")
	idx2 = strings.Index(s2, "</span>")
	sGao := s2[idx1+2 : idx2]

	sStart = "最　低："
	sEnd = "今　开："
	idx1 = strings.Index(s1, sStart)
	idx2 = strings.Index(s1, sEnd)
	s2 = s1[idx1+len(sStart) : idx2]
	idx1 = strings.Index(s2, "\">")
	idx2 = strings.Index(s2, "</span>")
	sDi := s2[idx1+2 : idx2]
	sStart = "今　开："
	sEnd = "昨　收："
	idx1 = strings.Index(s1, sStart)
	idx2 = strings.Index(s1, sEnd)
	s2 = s1[idx1+len(sStart) : idx2]
	idx1 = strings.Index(s2, "\">")
	idx2 = strings.Index(s2, "</span>")
	sKai := s2[idx1+2 : idx2]

	sStart = "成交额："
	sEnd = "成交量"
	idx1 = strings.Index(s1, sStart)
	idx2 = strings.Index(s1, sEnd)
	s2 = s1[idx1+len(sStart) : idx2]
	idx1 = strings.Index(s2, "</p>")
	s3 := s2[:idx1]
	idx1 = strings.Index(s3, "亿")
	sJE := s3[:idx1]

	sStart = "成交量："
	sEnd = "市盈率"
	idx1 = strings.Index(s1, sStart)
	idx2 = strings.Index(s1, sEnd)
	s2 = s1[idx1+len(sStart) : idx2]
	idx1 = strings.Index(s2, "</p>")
	s3 = s2[:idx1]
	idx1 = strings.Index(s3, "万")
	sJL := s3[:idx1]

	fEnd, err := strconv.ParseFloat(sCur, 64)
	fStart, err := strconv.ParseFloat(sKai, 64)
	fTradeFounds, err := strconv.ParseFloat(sJE, 64)
	fTradeStock, err := strconv.ParseFloat(sJL, 64)
	fHigh, err := strconv.ParseFloat(sGao, 64)
	fLow, err := strconv.ParseFloat(sDi, 64)
	logs.Debug("KAI:", sKai, "; SHOU:", sCur, ";GAO:", sGao, ";DI:", sDi, ";LANG:", sJL, ";E:", sJE, ";TIME:", dt)
	return Stock{stk, fHigh, fLow, fStart, fEnd, fTradeStock, fTradeFounds, dt}
	//	fmt.Sprintln("最高价：", sGao, "\n最低价：", sDi, "\n开盘价：", sKai, "\n当前价：", sCur, "\n成交额：", sJE, "\n成交量：", sJL)
}
