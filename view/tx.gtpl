{{ define "tx"}}
<div align="center" class="tx" >
<h3>腾讯视频</h3>
<a href="http://film.qq.com/weixin/upcoming.html"> http://film.qq.com/weixin/upcoming.html  -- 即将播出</a>
<table border="1" cellspacing="0" cellpadding="0" width=720>
<tr> <th>ID</th><th>Name</th><th>观演热度</th><th>发布时间</th><th>豆瓣分数</th></tr>
{{range .}}
{{if eq .Company "TX" }}
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
