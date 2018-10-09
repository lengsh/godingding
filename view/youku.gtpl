{{define "youku"}}
<div align="center" class="youku" >
<h3>优酷视频</h3>
<a href="https://vip.youku.com/vips/index.html">https://vip.youku.com/vips/index.html -- 即将上线 </a>
<table border="1" cellspacing="0" cellpadding="0" width=720>
<tr> <th>ID</th><th>Name</th><th>观演指数</th><th>ReleaseTime</th><th>豆瓣分数</th></tr>
{{range .}}
{{if eq .Company "YOUKU" }}
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
