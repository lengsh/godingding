{{template "header" .}}
<BR>
<div align="center">
<table><tr><td>
<div id="graph" >Loading graph...</div>
</td></tr></table>
</div>
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
<div align=center>
BABA,FB,MSFT,AMZN,AAPL,TSLA,BIDU,NVDA,GOOGL,WB
</div>
<script type="text/javascript">

var request = new XMLHttpRequest();
 request.open('GET', 'http://47.105.107.171/jsonp?func=stock&para=baba', true);
 request.onload = function () {
	 // Begin accessing JSON data here
	 var data = JSON.parse(this.response);
	 if (request.status >= 200 && request.status < 400) {
		 var myData = new Array();
		 idx = 1
		 data.forEach( stock => {
				 //                              console.log(stock.key);
				 //                              console.log(stock.value);
				 //  var vd = new Array([ parseInt(stock.key), parseInt(stock.value) ]);
				 // myData.push(vd[0]);
				 var vd = [ idx , parseInt(stock.value) ];
				 myData.push(vd);
				 idx++;
				 });

		 var myChart = new JSChart('graph', 'line');
		 myChart.setDataArray(myData);
		 myChart.setTitle('BABA Stock Trend');
		 myChart.setTitleColor('#8E8E8E');
		 myChart.setTitleFontSize(11);
		 myChart.setAxisNameX('');
		 myChart.setAxisNameY('');
		 myChart.setAxisColor('#C4C4C4');
		 myChart.setAxisValuesColor('#343434');
		 myChart.setAxisPaddingLeft(100);
		 myChart.setAxisPaddingRight(120);
		 myChart.setAxisPaddingTop(50);
		 myChart.setAxisPaddingBottom(40);
		 myChart.setAxisValuesNumberX(6);
		 myChart.setGraphExtend(true);
		 myChart.setGridColor('#c2c2c2');
		 myChart.setLineWidth(2);
		 myChart.setLineColor('#9F0505');
		 myChart.setSize(616, 321);
		 myChart.setBackgroundImage('chart_bg.jpg');
		 myChart.draw();

	 } else {
		 console.log('error');
	 }
 }
request.send();

</script>

{{ template "footer" .}}
