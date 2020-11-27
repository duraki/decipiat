## Templating (Engine)

### Include variable in loop

```
# => public/views/page.html
{{define "page"}}
	<code style="padding: 0px;">Welcome back.<br>Logged-in as <strong>{{ range $key, $value := .}} {{ $value.Email }} {{ end }}</strong>.</pre></code>
{{end}}
```

### Include variable (fixed)

```
# => partial
{{define "page"}}
	<br>
	<code style="padding: 0px;"><pre>Welcome back.<br>Logged-in as <strong>{{ . }}</strong>.</pre></code>
{{end}}

# => include
{{template "page" (index . "user").Email}}
```

