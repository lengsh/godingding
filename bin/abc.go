
package main

import(
 "net/http"
 "time"
 "io/ioutil"
 "fmt"
 "os"
 "strings"
 "log"
)


func main(){
 stk := "BABA"
 if len(os.Args) > 1  {
         stk = os.Args[1]
            stk = strings.ToUpper(stk)  
        }
  surl := fmt.Sprintf("https://www.futunn.com/quote/stock?m=us&code=%s",stk)

 tr := &http.Transport{
         MaxIdleConns:       10,
	 IdleConnTimeout:    30 * time.Second,
	 DisableCompression: true,
 }
 client := &http.Client{Transport: tr}
	resp, err := client.Get( surl )
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
/*
// <div class="time">
                idx1 = strings.Index(s1, "<div class=\"time\">")
                idx2 = strings.Index(s1, "<div class=\"stock_detail\">")
                s2 := s1[idx1+30:idx2]
                idx1 = strings.Index(s2, "\">") 
                idx2 = strings.Index(s2, "</span>")
                sTime := s2[idx1+2:idx2]
*/

                sStart := "<span class=\"price01\">"
                sEnd := "<script type=\"text/template\" id=\"basicQuoteTpl\">"
                idx1 := strings.Index(str, sStart)
                idx2 := strings.Index(str, sEnd)
                s1 := str[idx1+ len(sStart):idx2]
                sEnd = "</span>"
                idx2 = strings.Index(s1, sEnd)
                sCur := s1[0:idx2]

                sStart = "最　高："
                sEnd = "<p>最　低："
                idx1 = strings.Index(s1, sStart)
                idx2 = strings.Index(s1, sEnd)
                s2 := s1[idx1+ len(sStart):idx2]
                idx1 = strings.Index(s2, "\">")
                idx2 = strings.Index(s2, "</span>")
		sGao :=  s2[idx1+2:idx2]

                sStart = "最　低："
                sEnd = "今　开："
                idx1 = strings.Index(s1, sStart)
                idx2 = strings.Index(s1, sEnd)
                s2 = s1[idx1+ len(sStart):idx2]
                idx1 = strings.Index(s2, "\">")
                idx2 = strings.Index(s2, "</span>")
		sDi :=  s2[idx1+2:idx2]

                sStart = "今　开："
                sEnd = "昨　收："
                idx1 = strings.Index(s1, sStart)
                idx2 = strings.Index(s1, sEnd)
                s2 = s1[idx1+ len(sStart):idx2]
                idx1 = strings.Index(s2, "\">")
                idx2 = strings.Index(s2, "</span>")
		sKai :=  s2[idx1+2:idx2]

                sStart = "成交额："
                sEnd = "成交量"
                idx1 = strings.Index(s1, sStart)
                idx2 = strings.Index(s1, sEnd)
                s2 = s1[idx1+ len(sStart):idx2]
                idx1 = strings.Index(s2, "</p>")
		sJE :=  s2[:idx1]

                sStart = "成交量："
                sEnd = "市盈率"
                idx1 = strings.Index(s1, sStart)
                idx2 = strings.Index(s1, sEnd)
                s2 = s1[idx1+ len(sStart):idx2]
                idx1 = strings.Index(s2, "</p>")
		sJL :=  s2[:idx1]


fmt.Println("最高价：", sGao,"\n最低价：", sDi,"\n开盘价：", sKai,"\n当前价：", sCur,"\n成交额：",sJE,"\n成交量：",sJL)

}
