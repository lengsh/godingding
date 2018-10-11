<html>
<div align="center" width=520>
<table border="1" cellspacing="0" cellpadding="0">
<tr> <th>ID</th><th width=80>Name</th><th width=80>最高</th><th width=80>最低</th><th width=80>开盘</th><th width=80>收盘</th><th width=80>交易额</th><th width=80>交易量</th><th width=80>市值</th><th width=120>Date</th></tr>
{{range .}}

<tr><th>{{.Id}} </th><th>{{.Name}} </th><th> {{.HighPrice}} </th><th>{{.LowPrice}}</th><th>{{.StartPrice}}</th><th>{{.EndPrice}}</th><th>{{.TradeFounds}}</th><th>{{.TradeStock}}</th><th>{{.MarketCap}}</th><th>{{.CreateDate}}</th> </tr>

{{end}}

</table>
</div>

<BR>
<div align="center">

<TR><a href="/movie/first">Home</a><TR>
<TR><a href="/query?do=stock">Stock</a><TR>
<TR><a href="/query?do=movie">Movie</a><TR>
<TR><a href="/query?do=report">Report</a><TR>
</div>
</html>
