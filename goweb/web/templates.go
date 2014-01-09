package web

const (
	ADMIN_HOME=`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
<title>administrator</title>
</head>
<body>
<h3>welcome!</h3>
<table>
<tr><td bgcolor="#eaeaea">
<form action="./" method="POST" enctype="multipart/form-data">
	tarfile:<input type="file" name="filename" /><input type="submit" />
	<input type="hidden" name="action" value="tarupload" />
</form>
</td></tr>
<tr><td bgcolor="#eaeaea">
<form action="./" method="POST" enctype="multipart/form-data">
	zipfile:<input type="file" name="filename" /><input type="submit" />
	<input type="hidden" name="action" value="zipupload" />
</form>
</td></tr>
<tr><td bgcolor="#eaeaea">
<form action="./" method="POST" enctype="multipart/form-data">
	filepath:<input name="filepath" />
	file:<input type="file" name="filename" />
	<input type="submit" />
	<input type="hidden" name="action" value="upload" />
</form>
</td></tr>
</table>
<a href="click">statistic</a>
<a href="template">template</a>
<a href="list">file list</a>
<a href="bloblist">blob list</a>
<a href="listsession">session list</a>
</body>
</html>
`
        TMPL_ADMIN=` <!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
        <title>upload templates</title>
</head>
<body>
<a href="./">admin</a>

<form action="template" method="POST"enctype="multipart/form-data">
<select name="tmpl_name">{{range $idx,$elem := .tlist}}<option value="{{$idx}}">{{$elem}}</option>{{end}}</select>
<input type="file" name="filename" /> <input type="submit"/>
</form>
</body>
</html>
`
	LIST=`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
	<title>list file</title>
</head>
<body>
<a href="./">admin</a>
{{if .data}}
<table><tr><td>Name</td><td>Update</td></tr>
{{end}}
{{range .data}}
<tr><td><a href="bykey?bk={{.Key}}">{{.Name}}</a></td><td>{{.Update}}</td></tr><td></td>{{end}}
{{if .data}}
</table>
{{end}}

</body>
</html>
`
	BLOBLIST=`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
	<title>list file</title>
</head>
<body>
<a href="./">admin</a>
{{if .data}}
<table><tr><td>Name</td><td>Update</td></tr>
{{end}}
{{range .data}}
<tr><td>{{.Md5str}}</td><td>{{.Update}}</td></tr><td></td>{{end}}
{{if .data}}
</table>
{{end}}

</body>
</html>
`
	CLICKLIST=`{{$f2 :=.int64time }}{{$f1 := .odd }}{{$add := .add }}{{$start := .start}}
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
	<title>click history</title>
</head>
<body>
<a href="./">admin</a>
Total: {{.total}}<a href="click.{{.lastpage}}">last page</a>
{{if .data}}
<table><tr><td></td><td>IP</td><td>Agent</td><td>Request</td><td>When</td></tr>
{{end}}
{{range $idx,$v := .data}}
{{if call $f1 $idx}}
<tr bgcolor="#eaeaea"><td>{{call $add $start $idx}}</td><td>{{.Ip}}</td><td>{{.Agent}}</td><td>{{.Request}}</td><td>{{call $f2 .When}}</td></tr>
{{else}}
<tr bgcolor="#aeaeae"><td>{{call $add $start $idx}}</td><td>{{.Ip}}</td><td>{{.Agent}}</td><td>{{.Request}}</td><td>{{call $f2 .When}}</td></tr>
{{end}}
{{end}}
{{if .data}}
</table>
{{end}}

</body>
</html>
`
	SESSIONLIST=`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
	<title>list file</title>
</head>
<body>
<a href="./">admin</a>
{{if .data}}
<table><tr><td>Name</td><td>Update</td></tr>
{{end}}
{{range $k,$v := .data}}
<tr><td>{{$k}}</td><td>{{$v}}</td></tr>{{end}}
{{if .data}}
</table>
{{end}}

</body>
</html>
`)
