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

var TIMELEN int64 = 180 // 60s 有效时长

func createScrumb(delay int64) string {
	//  usec := time.Now().UnixNano()
	//  secs  := usec/(1000000*1000)
	secs := time.Now().Unix()
	pnum := secs/(TIMELEN/2) - delay

	str := fmt.Sprintf("%s%d Scrumb secret keY", ServerConfig.PassSalt, pnum)
	h := md5.New() // 计算MD5散列，引入crypto/md5 并使用 md5.New()方法。
	h.Write([]byte(str))
	bs := h.Sum(nil)
	//      log.Println(pnum)
	//ss := fmt.Sprintf("%x", bs)
	//   fmt.Println(pnum, "; create, salt=", ss)
	return fmt.Sprintf("%x", bs)
}

func CreateScrumb() string {
	return createScrumb(0)
}

func CheckScrumb(old string) bool {
	s1 := createScrumb(0)
	if s1 == old {
		return true
	}
	s1 = createScrumb(1)
	if s1 == old {
		return true
	}
	return false
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
