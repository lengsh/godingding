{{define "iqiyi"}}
<div align="center" class="iqiyi" >
<h3>爱奇艺</h3>
<a href="http://m.iqiyi.com/vip/timeLine.html"> http://m.iqiyi.com/vip/timeLine.html -- 即将上线(移动浏览器) </a>
<table border="1" cellspacing="0" cellpadding="0" width=720>
<tr> <th>ID</th><th>Name</th><th>订阅人数</th><th>发布时间</th><th>豆瓣分数</th></tr>
{{range .}}
{{if eq .Company "IQIYI" }}
<tr><th>{{.Id}} </th><th align="left">{{.Name}} </th><th align="left">  
{{if .IfTagRate }}
<font color="RED">{{.Rate}} </font>
{{else}}
     {{.Rate}}
{{end}}
</th><th>{{.Releasetime}}</th><th>
{{if (ge .Douban 8.0) }}
<font color="RED">{{.Douban}} </font>
{{else if (lt .Douban 5.0)}}
<font color="GREEN">{{.Douban}} </font>
{{else}}
{{.Douban}}
{{end}}
</th> </tr>
{{end}}
{{end}}
</table>
</div>
{{end}}
