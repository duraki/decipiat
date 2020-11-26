package web

import (
	"github.com/duraki/decipiat/web/handlers"
	"github.com/labstack/echo"
	_ "github.com/labstack/echo/v4/middleware"
	_ "html/template"
	_ "io"
	"log"
)

/*
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
*/

/* sample -- https://github.com/xesina/golang-echo-realworld-example-app/tree/master/router */
func Init() *echo.Echo {
	tmpl, _ := NewTmpl("public/views/", "html", true)
	err := tmpl.Load()
	if err != nil {
		log.Printf("err => %s", err.Error)
	}

	// t := &Template{
	// 	templates: template.Must(template.ParseGlob("public/views/**/*")),
	// }

	e := echo.New()
	e.Renderer = tmpl

	adminGroup := e.Group("/admin") /* create groups */

	MainGroup(e)
	AdminGroup(adminGroup)

	return e
}

func MainGroup(e *echo.Echo) {
	// Route / to handle defaults
	e.GET("/", handlers.Homepage)
	e.GET("/status", handlers.Status)

	/*
		e.GET("/", handler.Home)
		e.GET("/health-check", handler.HealthCheck)

		e.GET("/user/register", handler.RegisterUserView)
		e.POST("/user/register", handler.RegisterUser)

		e.GET("/user/login", handler.LoginUserView)
		e.POST("/user/login", handler.LoginUserView)
	*/

	//e.GET("/project/create", handlers.ProjectCreateView)
	//e.POST("/project/create", handlers.ProjectCreate)
}

func UserGroup(g *echo.Group) {
	g.GET("/project/create", handlers.ProjectCreate)
}

func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
