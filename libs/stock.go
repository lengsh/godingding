
 package  libs 
    
  import (
  "github.com/benbjohnson/phantomjs"
   "fmt"
   // "os"
   "strings"
)


func Crawler_Phantomjs( st string) string{

if err := phantomjs.DefaultProcess.Open(); err != nil {
	fmt.Println(err)
		return ""
}
defer phantomjs.DefaultProcess.Close()
page,err := phantomjs.CreateWebPage()
	if err != nil {
		return ""
	}
	stk := strings.ToUpper(st)  
surl := fmt.Sprintf("http://quotes.money.163.com/usstock/%s.html#US1a01",stk)

if err := page.Open(surl); err != nil {
	return ""
}

   if content,err := page.Content();err == nil{
         idx1 :=  strings.Index(content, "<div class=\"stock_info\">")
         idx2 :=  strings.Index(content,"<div class=\"stock_nav_bar\">")
	 s1 := content[idx1:idx2]

         idx1 = strings.Index(s1, "<div class=\"time\">")
         idx2 = strings.Index(s1, "<div class=\"stock_detail\">")
         s2 := s1[idx1+30:idx2]
         idx1 = strings.Index(s2, "\">") 
         idx2 = strings.Index(s2, "</span>")
         sTime := s2[idx1+2:idx2] 
                  
    return strings.ToUpper(sTime)
       }
       return ""
 }
