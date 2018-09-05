
package main

import(
 "fmt"
 "os"
 "strings"
 "github.com/lengsh/godingding/libs"
)


func main(){
 stk := "BABA"
 if len(os.Args) > 1  {
         stk = os.Args[1]
            stk = strings.ToUpper(stk)  
        }
 s := libs.Crawler_Futu(stk)
 fmt.Println(s)

}
