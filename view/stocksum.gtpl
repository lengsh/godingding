{{template "header" .}}

<div align="center" width=520>
<table border="1" cellspacing="0" cellpadding="0">
<tr> <th width=80>市值</th><th width=120>Date</th></tr>
{{range .}}
<tr><th>

{{if (ge .SumMarket 66000.0) }}
 <font color="RED">{{.SumMarket}} </font>
{{else if (lt .SumMarket 46000.0)}}
 <font color="GREEN">{{.SumMarket}} </font>
{{else}}
{{.SumMarket}}
{{end}}

</th><th>{{.TString}}</th> </tr>
{{end}}

</table>
</div>
<BR>

{{template "footer"}}
