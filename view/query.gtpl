<html>
<br><br><BR>
<div align="center">
<table border="1" >
<tr> <th>ID</th><th>Name</th><th>Department</th><th>Number</th></tr>

<form action="login?"  >
{{range .}}
<tr><th>{{.Id}} </th><th>{{.Name}} </th><th> {{.Department}} </th><th>{{.Number}}</th></tr>
{{end}}

</table>
</div>

</html>
