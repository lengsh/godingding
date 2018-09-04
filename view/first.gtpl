<html>
<br><br><BR>
<div align="center">

<image src="https://alinw-oss.alicdn.com/alinw-node-admin-public-oss/2018-7-31/1533034537081/280X180.jpg?x-oss-process=image/resize,m_fixed,h_360,w_560" />

<a href="/send?">Test</a><TR>

<table border="1" >
<tr> <th>ID</th><th>Name</th><th>Department</th><th>Number</th></tr>
{{range .}}
<tr><th>{{.Id}} </th><th>{{.Name}} </th><th> {{.Department}} </th><th>{{.Number}}</th></tr>
{{end}}
</table>
</div>

</html>
