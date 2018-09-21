<html>
<br><br><BR>
<div align="center" width=450 >
<table border="1" cellspacing="0" cellpadding="0">
<tr> <th>ID</th><th>Name</th><th>Rate</th><th>ReleaseTime</th><th>Company</th></tr>

{{range .}}
<tr><th>{{.Id}} </th><th align="left">{{.Name}} </th><th> {{.Rate}} 万 </th><th>{{.Releasetime}}</th><th>{{.Company}}</th> </tr>
{{end}}

</table>
</div>

<div align="center">
<TR><a href="/send?">Test</a><TR>
</div>
</html>
