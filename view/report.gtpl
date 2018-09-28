<html>
<div class="updatetime" align=right>
{{ (index . 0).UpdateTime }} 
</div>
<div class="TX" align="center" >
<h3>腾讯视频</h3>
<table border="1" cellspacing="0" cellpadding="0" width=720>
<tr> <th>ID</th><th>Name</th><th>热度指数</th><th>ReleaseTime</th><th>Company</th><th>豆瓣分数</th></tr>

{{range .}}
{{if eq .Company "TX" }}
<tr><th>{{.Id}} </th><th align="left">{{.Name}} </th><th align="left">  

{{if .IfTagRate }}
<font color="RED">{{.Rate}} </font>
{{else}}
     {{.Rate}}
{{end}}
</th><th>{{.Releasetime}}</th><th>{{.Company}}</th><th>
{{if (ge .Douban 8.0) }}
<font color="RED">{{.Douban}} </font>
{{else}}
{{.Douban}}
{{end}}
</th> </tr>
{{end}}
{{end}}
</table>
</div>

<div align="center" class="iqiyi" >
<h3>爱奇艺</h3>
<table border="1" cellspacing="0" cellpadding="0" width=720>
<tr> <th>ID</th><th>Name</th><th>订阅人数</th><th>ReleaseTime</th><th>Company</th><th>豆瓣分数</th></tr>
{{range .}}
{{if eq .Company "IQIYI" }}
<tr><th>{{.Id}} </th><th align="left">{{.Name}} </th><th align="left">  
{{if .IfTagRate }}
<font color="RED">{{.Rate}} </font>
{{else}}
     {{.Rate}}
{{end}}
</th><th>{{.Releasetime}}</th><th>{{.Company}}</th><th>
{{if (ge .Douban 8.0) }}
<font color="RED">{{.Douban}} </font>
{{else}}
{{.Douban}}
{{end}}
</th> </tr>
{{end}}
{{end}}
</table>
</div>

<div align="center" class="youku" >
<H3>优酷视频</H3>
<table border="1" cellspacing="0" cellpadding="0" width=720>
<tr> <th>ID</th><th>Name</th><th>影片指数</th><th>ReleaseTime</th><th>Company</th><th>豆瓣分数</th></tr>
{{range .}}
{{if eq .Company "YOUKU" }}
<tr><th>{{.Id}} </th><th align="left">{{.Name}} </th><th align="left">  
{{if .IfTagRate }}
<font color="RED">{{.Rate}} </font>
{{else}}
     {{.Rate}}
{{end}}
</th><th>{{.Releasetime}}</th><th>{{.Company}}</th><th>
{{if (ge .Douban 8.0) }}
<font color="RED">{{.Douban}} </font>
{{else}}
{{.Douban}}
{{end}}
</th> </tr>
{{end}}
{{end}}
</table>
</div>
<BR><BR>
<div align="center">
</div>
</html>
