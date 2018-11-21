{{template "header" .}}
<BR>
<div align="center">
 <table><tr><td>
 <div id="graph" >Loading graph...</div>
 </td></tr></table>
 </div>

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

<script type="text/javascript">
var request = new XMLHttpRequest();
request.open('GET', 'http://47.105.107.171/jsonp?func=stocksum', true);
request.onload = function () {
	// Begin accessing JSON data here
	var data = JSON.parse(this.response);
	if (request.status >= 200 && request.status < 400) {
        var myData = new Array();
	idx = 1;
var myChart = new JSChart('graph', 'line');

data.forEach( stock => {
//				console.log(stock.key);
//				console.log(stock.value);
//  var vd = new Array([ parseInt(stock.key), parseInt(stock.value) ]);
// myData.push(vd[0]);
		var vd = [idx, parseInt(stock.value) ];
		myData.push(vd);
		if ((idx%2==0) && (idx != 0)) {
		myChart.setLabelX([idx, stock.key]);
		}
		myChart.setTooltip([idx]);
		idx++;
		});

// myData = [[1, 100],[2, 555],[3, 80],[4, 115],[5, 580],[6, 70],[7, 30],[8, 130],[9, 160],[10, 170]];
myChart.setDataArray(myData, 'red');
myChart.setAxisPaddingBottom(40);
myChart.setTextPaddingBottom(10);
myChart.setAxisValuesNumberY(6);
myChart.setIntervalStartY(350);
myChart.setIntervalEndY(520);

myChart.setAxisValuesNumberX(8);
myChart.setShowXValues(false);
myChart.setTitleColor('#454545');
myChart.setAxisValuesColor('#454545');
myChart.setLineColor('#FF0101', 'red');

    myChart.setTitle('10 Stock Trend');
                       myChart.setTitleColor('#8E8E8E');
                       myChart.setTitleFontSize(11);

myChart.setFlagColor('#9D16FC');
myChart.setFlagRadius(4);
myChart.setBackgroundImage('chart_bg.jpg');
myChart.setSize(616, 321);
myChart.draw();
} else {
		console.log('error');
	}
}
request.send();

</script>

{{template "footer"}}
