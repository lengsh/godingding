<html>
<br><br><BR>
<div align="center">
<image src="https://is.golangtc.com/logo/golangtc.png?height=160" />
</div>
<BR>
<div align=center>

<table border="0" cellspacing="0" cellpadding="0">
 {{range .}}
 <tr><th align="left" width=200>{{.Name}} </th><th align="left" width=80>
{{ if .IfTagRate}}
 <font color="RED">{{.Rate}} </font>
 {{else}}
 {{.Rate}}
 {{end}}
 </th><th width=120>{{.Releasetime}}</th><th width=120>{{.Company}}</th> 
 </th><th width=100>
{{if (ge .Douban 8.0)}}
<font color=RED>{{.Douban}}</font>
{{else}}
   {{if (lt .Douban 5.0)}}
<font color=GREEN>{{.Douban}}</font>
   {{else}}
{{.Douban}}
{{end}}
{{end}}
</th> </tr>
 {{end}}
 </table>
</div>

<BR>
<div align="center" width=300>
<TR><a href="/query?do=stock">Stock</a>&nbsp;&nbsp;&nbsp;<TR>
<TR><a href="/query?do=report">Report</a>&nbsp;&nbsp;&nbsp;<TR>
<TR><a href="/query?do=movie">Movie</a><TR>
</div>

</html>
