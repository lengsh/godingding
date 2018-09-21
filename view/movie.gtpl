<html>
<div align="center" width=450 >
<table border="1" cellspacing="0" cellpadding="0">
<tr> <th>ID</th><th>Name</th><th>Rate</th><th>ReleaseTime</th><th>Company</th></tr>

{{range .}}
<tr><th>{{.Id}} </th><th align="left">{{.Name}} </th><th align="left">  
{{ if .TagRate}}
<font color="RED">{{.Rate}} </font>
{{else}}
{{.Rate}}
{{end}}
</th><th>{{.Releasetime}}</th><th>{{.Company}}</th> </tr>
{{end}}

</table>
</div>
<BR><BR>
<div align="center">
<TR><a href="/first?">Home</a><TR>
<TR><a href="/send?">Test</a><TR>
<TR><a href="/query?do=stock">Stock</a><TR>
<TR><a href="/query?do=movie">Movie</a><TR>

</div>
</html>
