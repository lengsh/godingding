{{template "header" .}}
<BR>
<div align="center" >
<table border="1" cellspacing="0" cellpadding="0">
<tr> <th width=50>ID</th><th width=80>Name</th><th width=80>最高</th><th width=80>最低</th><th width=80>开盘</th><th width=80>收盘</th><th width=80>交易额</th><th width=80>交易量</th><th width=80>市值</th><th width=120>Date</th></tr>
{{range .}}

<tr><th>{{.Id}} </th><th><a href="http://47.105.107.171/stock/stock?do={{.Name}}">{{.Name}}</a></th><th> {{.HighPrice}} </th><th>{{.LowPrice}}</th><th>{{.StartPrice}}</th><th>{{.EndPrice}}</th><th>{{.TradeFounds}}</th><th>{{.TradeStock}}</th><th>{{.MarketCap}}</th><th>{{.TString}}</th> </tr>

{{end}}

</table>
</div>

<BR>
<div align=center>
<h3>10标总市值：<font color=red>{{ (index . 0).SumMarketCap }} </font> </H3>
<a href="http://47.105.107.171/stock/stock?s=all"> 趋势&历史 </a> <BR>
</div>

{{ template "footer" .}}
