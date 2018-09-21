<html>
<BR>
<div align=center> 
 <TR><a href="/first">HomePage</a><TR>
<TR> <TR><a href="/send?">SendMessage</a><TR>
<TR> <TR><a href="/query?do=stock">Stock</a><TR>
<TR> <TR><a href="/query?do=movie">Movie</a><TR>
</div>
<BR><BR>
<div align="center">
<table border="0">
<form action="login?"  >
<tr>
<th>Message: </th><th><input name="message" type=text size=48 /> </th>
</tr>
<tr>
<th><input name=".scrumb" type=hidden value="{{.Scrumb}}" /></th><th></th>
</tr>
<tr><th></th><th>
<input type="submit" value="Send YY" formaction="/send?" />&nbsp;&nbsp;&nbsp;&nbsp;
<input type="submit" value="Send MH" formaction="/query?" />
</th>
</tr>
</table>
</div>
</html>
