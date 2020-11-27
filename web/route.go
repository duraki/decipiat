package web

import (
	"github.com/duraki/decipiat/web/handlers"

	// "github.com/duraki/decipiat/web/session"
	// "github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	_ "html/template"
	_ "io"
)

/*
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
*/

func SetGlobals() {
	db, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("%s\n", "decipiat starting a db conn to ... localhost")
	}

	// Create indices
	if err = db.Copy().DB("decipiat").C("users").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatalf("%s\n", "unable to create db indices ...")
	}

	// Initialize a global handler
	handlers.GlobalConfig = handlers.Globals{DB: db}
}

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
	//e.Use(session.Middleware(Sessions.NewCookieStore([]byte("secret"))))

	SetGlobals()
	log.Infof("global handlers %+v\n", handlers.GlobalConfig)

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
	e.POST("/register", handlers.RegisterUser)
	e.GET("/login", handlers.LoginUserView)
	e.POST("/login", handlers.LoginUser)
	e.GET("/logout", handlers.LogoutUser)
	e.GET("/me", handlers.UserDashboardView)

	// Route for Project Management
	e.GET("/project/new", handlers.ProjectCreateView)
	e.POST("/project/new", handlers.ProjectCreate)

	// Route for domain generation
	e.GET("/domain", handlers.DomainView)
	e.POST("/domain", handlers.SearchDomain)

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

func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
