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

## MongoDB

### Find All Documents (Document == ROW)

```
func ProjectListView(c echo.Context) (err error) {
	db := GlobalConfig.DB.Clone()

	var projects []*models.Project

	defer db.Close()
	if err = db.DB(DatabaseName).C(models.CollectionProject).Find(nil).All(&projects); err != nil {
		fmt.Errorf("%s %+v", "Error while retrieve Project List View, User", session.GetUser())
	}

	// do something w/ the `&projects`
	...
```

### Find Specific Document

```

```

Resources:

	* https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
