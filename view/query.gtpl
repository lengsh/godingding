<html>
<br><br><BR>
<div align="center">
<table border="1" >
<tr> <th>ID</th><th>Name</th><th>High</th><th>Low</th><th>Start</th><th>Current</th><th>Founds</th><th>Stocks</th><th>Date</th></tr>

<form action="login?"  >
{{range .}}
<tr><th>{{.Id}} </th><th>{{.Name}} </th><th> {{.HighPrice}} </th><th>{{.LowPrice}}</th><th>{{.TradeStart}}</th><th>{{.TradeEnd}}</th><th>{{.TradeFounds}}</th><th>{{.TradeStock}}</th><th>{{.TradeDate}}</th> </tr>
{{end}}

</table>
</div>

<div align="center">
<TR><a href="/send?">Test</a><TR>
</div>
</html>
