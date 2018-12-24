package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func GetCaller(depth int) string {
	_, file, line, ok := runtime.Caller(depth)
	if !ok {
		file = "???"
		line = 0
	}
	return fmt.Sprintf("%s,%d", file, line)
}

func RecoverDefer(position string) {
	if r := recover(); r != nil {
		dingtalker := NewDingtalker()
		s := fmt.Sprintf("[%s] %s", position, r)
		dingtalker.SendChatTextMessage(s)
		logs.Error(r)
	}
}

func CreateScrumb(s string) string {
	secs := time.Now().Unix()
	pnum := secs / 60
	str := fmt.Sprintf("%s%d Scrumb secret keY", s, pnum)
	h := md5.New() // 计算MD5散列，引入crypto/md5 并使用 md5.New()方法。
	h.Write([]byte(str))
	bs := h.Sum(nil)
	//      log.Println(pnum)
	//      log.Println(fmt.Sprintf("%x", bs))
	return fmt.Sprintf("%x", bs)
}

func AuthenticationIP(req *http.Request) (error, bool) {
	//	sip = "0:0:1;127.0.0.1;47.105.107.171"
	//	svp := strings.Split(sip, ";")
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return err, false
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		//return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
		fmt.Printf("userip: %q is not IP:port", req.RemoteAddr)
		return err, false
	}
	return nil, true
}

func PullwordGet(s string, param1 int, param2 int, level float64) (map[string]float64, bool) {
	retv := map[string]float64{}
	url := fmt.Sprintf("http://api.pullword.com/get.php?source=%s&param1=%d&param2=%d", s, param1, param2)
	resp, err := http.Get(url)
	if err != nil {
		logs.Debug(err)
		return retv, false
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Debug(err)
		return retv, false
	}
	sv := strings.Split(string(body), "\r\n")
	for _, v := range sv {
		if len(v) > 3 {
			ws := strings.Split(v, ":")
			if len(ws) == 2 {
				word := ws[0]
				prob, err := strconv.ParseFloat(ws[1], 32)
				if err != nil {
					logs.Debug(err)
					prob = 0
				}

				if prob > level {
					retv[word] = prob
					//fmt.Println(word, ":", prob, "; ", len(word))
				}
			}
		}
	}
	return retv, true
}
