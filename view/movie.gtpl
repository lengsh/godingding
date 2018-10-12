{{ template "header" . }}
<div align="center" width=450 >
<table border="1" cellspacing="0" cellpadding="0">
<tr> <th>ID</th><th>Name</th><th>Rate</th><th>ReleaseTime</th><th>Company</th><th>Douban</th></tr>

{{range .}}
{{$name := .Name }}
<tr><th>{{.Id}} </th><th align="left">{{.Name}} </th><th align="left">  

{{if .IfTagRate }}
<font color="RED">{{.Rate}} </font>
{{else}}
     {{.Rate}}
{{end}}

</th><th>{{.Releasetime}}</th><th>{{.Company}}</th><th>

{{if (ge .Douban 8.0) }}
<font color="RED">{{.Douban}} </font>
{{else if (lt .Douban 5.0)}}
<font color="GREEN">{{.Douban}} </font>
{{else}}
{{.Douban}}
{{end}}

</th> </tr>

{{end}}
</table>
</div>

{{ template "footer" . }}
