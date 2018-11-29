package libs

import (
	"crypto/md5"
	"fmt"
	"net"
	"net/http"
	"time"
)

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
