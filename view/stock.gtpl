<html>
<div align="center" width=520>
<table border="1" cellspacing="0" cellpadding="0">
<tr> <th>ID</th><th>Name</th><th>High</th><th>Low</th><th>Start</th><th>Current</th><th>Founds</th><th>Stocks</th><th>Date</th></tr>

<form action="login?"  >
{{range .}}
<tr><th>{{.Id}} </th><th>{{.Name}} </th><th> {{.HighPrice}} </th><th>{{.LowPrice}}</th><th>{{.StartPrice}}</th><th>{{.EndPrice}}</th><th>{{.TradeFounds}}</th><th>{{.TradeStock}}</th><th>{{.TradeDate}}</th> </tr>
{{end}}

</table>
</div>


<BR>
<div align="center">

<TR><a href="/first">Home</a><TR>
<TR><a href="/query?do=stock">Stock</a><TR>
<TR><a href="/query?do=movie">Movie</a><TR>
<TR><a href="/query?do=report">Report</a><TR>
</div>
</html>
