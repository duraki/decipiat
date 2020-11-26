package web

import (
	"github.com/duraki/decipiat/web/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	/**
	 * Load templates.
	 * @type {[type]}
	 */
	tmpl, _ := NewTmpl("public/views/", ".html", true)
	err := tmpl.Load()
	if err != nil {
		log.Printf("err => %s", err.Error)
	}

	/** @type {echo} new echo server */
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	/**
	 * Setup static files.
	 * Do not edit!
	 */
	e.Static("/", "public/views")

	/**
	 * Setup templating engine.
	 * @type {[type]}
	 */
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

	// Route for User Management
	e.GET("/register", handlers.RegisterUserView)

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
