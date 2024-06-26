package web

import (
	"github.com/duraki/decipiat/web/handlers"
	"github.com/duraki/decipiat/web/handlers/project"
	md "github.com/duraki/decipiat/web/middlewares"
	"github.com/duraki/decipiat/web/tfunctions"

	// "github.com/gorilla/sessions"
	"html/template"
	_ "html/template"
	_ "io"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

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

	// Initialize function map
	funcs := template.FuncMap{
		"percentage": tfunctions.Percent,
	}

	/**
	 * Load templates.
	 * @type {[type]}
	 */
	tmpl, _ := NewTmpl("public/views/", ".html", true, funcs)
	tmpl.Funcs(funcs)
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
	projectGroup := e.Group("/project")

	MainGroup(e)
	AdminGroup(adminGroup)
	ProjectGroup(projectGroup)

	return e
}

func ProjectGroup(g *echo.Group) {
	g.Use(md.IsAuthenticated)
	// Route for Project Management
	g.GET("/new", handlers.ProjectCreateView)
	g.POST("/new", handlers.ProjectCreate)
	g.GET("/list", handlers.ProjectListView)
	g.GET("/view/:cpvUuid", handlers.ProjectView)

	g.GET("/edit/:cpvUuid/domain", project.TargetElementariesView)
	g.POST("/edit/:cpvUuid/domain", project.TargetElementaries)
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

	// Route for domain generation
	e.GET("/domain", handlers.DomainView)
	e.POST("/domain", handlers.SearchDomain)
}

func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
