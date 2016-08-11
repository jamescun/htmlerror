package htmlerror

import (
	"html/template"
)

var tpl = template.Must(template.New("error").Parse(`<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />

		<title>{{ .Name }}: {{ .Error }}</title>

<style type="text/css">
html { margin: 0; padding: 0; }
body {
	margin: 0; padding: 0;

	color: #333333;
	background: #FFFFFF;
	font: 16px "Helvetica Neue", Helvetica, Arial, sans-serif;
}

h1, h2, h3,
h4, h5, h6 { margin: 0 0 25px 0; padding: 0; }

table {
	width: 100%;
	margin-bottom: 25px;

	text-align: left;
	font-size: 0.8em;
	border-spacing: 0;
}
table th, table td { padding: 10px; }
table th { font-weight: bold; text-transform: uppercase; }
table thead th { border-bottom: 2px solid #EDEDED; }
table tbody th { border-right: 2px solid #EDEDED; }

main { padding: 50px; }

main.error { background: #F9F9F9; border-bottom: 1px solid #EDEDED; }
main.error header { color: #375EAB; }
main.error header h3 { margin-bottom: 10px; }
main.error header h3 small { color: #666666; font-weight: normal; }
main.error header h1 { margin-bottom: 50px; font-weight: normal; }

main.error article.stacktrace { box-shadow: 0 0 5px rgba(0,0,0,0.1); }
main.error article.stacktrace section.frame {
	padding: 15px 10px;

	line-height: 1.5em;
	background: #FFFFFF;
	border-bottom: 1px solid #EDEDED;
}
main.error article.stacktrace section.frame:last-child { border-bottom: none; }
main.error article.stacktrace section.frame small { color: #666666; }
main.error article.stacktrace section.frame em { color: #666666; font-size: 0.9em; font-style: normal; }
main.error article.stacktrace section.frame pre {
	margin: 10px 0 0 0; padding: 5px;

	tab-size: 4;
	background: #F9F9F9;
	overflow-x: scroll;
}

main.meta article td { font-family: monospace; }
main.meta article label { display: inline-block; margin-bottom: 25px; font-weight: bold; }

.expand input[name="expand"] { display: none; }
.expand label {  cursor: pointer; }
.expand label small { font-weight: normal; }
.expand section { display: none; }
.expand input[name="expand"]:checked ~ section { display: block; }
</style>
	</head>
	<body>
		<main class="error">
			<header>
				<h3 class="type">{{ .Name }} <small class="request">at <code>{{ .Request.Method }} {{ .Request.URL.Path }} {{ .Request.Proto }}</code></small></h3>
				<h1 class="message">{{ .Error }}</h1>
			</header>

			<article class="stacktrace">
{{ range .Stacktrace }}
				<section class="frame">
					<strong>{{ .Function }}</strong> <small>{{ .Module }}</small><br/>
					<em>{{ .Filename }}:{{ .Line }}</em>
<pre>{{ .Context }}</pre>
				</section>
{{ end }}
			</article>
		</main>

		<main class="meta">
{{ if .Request.Header }}
			<article class="headers expand">
				<input type="checkbox" name="expand" id="headers" />
				<label for="headers">Request Headers <small>&rarr;</small></label>

				<section>
					<table class="headers">
						<thead>
							<tr>
								<th>Name</th>
								<th>Value</th>
							</tr>
						</thead>
						<tbody>
{{ range $key, $value := .Request.Header }}
							<tr>
								<td>{{ $key }}</td>
								<td>{{ range $value }}{{ . }}</br>{{ end }}
							</tr>
{{ end }}
						</tbody>
					</table>
				</section>
			</article>
{{ end }}

{{ if .Request.URL }}
			<article class="query expand">
				<input type="checkbox" name="expand" id="query" />
				<label for="query">Request Query Parameters <small>&rarr;</small></label>

				<section>
					<table class="query">
						<thead>
							<tr>
								<th>Name</th>
								<th>Value</th>
							</tr>
						</thead>
						<tbody>
{{ range $key, $value := .Request.URL.Query }}
							<tr>
								<td>{{ $key }}</td>
								<td>{{ range $value }}{{ . }}</br>{{ end }}
							</tr>
{{ end }}
						</tbody>
					</table>
				</section>
			</article>
{{ end }}

{{ if .Request.Form }}
			<article class="form expand">
				<input type="checkbox" name="expand" id="form" />
				<label for="form">Request Form <small>&rarr;</small></label>

				<section>
					<table class="query">
						<thead>
							<tr>
								<th>Name</th>
								<th>Value</th>
							</tr>
						</thead>
						<tbody>
{{ range $key, $value := .Request.Form }}
							<tr>
								<td>{{ $key }}</td>
								<td>{{ range $value }}{{ . }}</br>{{ end }}
							</tr>
{{ end }}
						</tbody>
					</table>
				</section>
			</article>
{{ end }}
		</main>
	</body>
</html>`))
