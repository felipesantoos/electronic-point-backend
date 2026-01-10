package main

import (
	"eletronic_point/src/apps/api/config"
	"eletronic_point/src/apps/api/middlewares"
	"eletronic_point/src/apps/api/routes"
	"eletronic_point/src/apps/api/views"
	"eletronic_point/src/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "eletronic_point/src/apps/api/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	godotenv.Load(".env")
	cfg := config.Load()
	setUpAuth(cfg.Authorization)
	api := NewAPI(getServerHostAndPort())
	api.Serve()
}

func getServerHostAndPort() (string, int) {
	host := utils.GetenvWithDefault("SERVER_HOST", "0.0.0.0")
	portStr := utils.GetenvWithDefault("SERVER_PORT", "8000")
	var port int
	if v, err := strconv.Atoi(portStr); err != nil {
		log.Fatal("The server port env variable must be a number (e.g 8000)")
	} else {
		port = v
	}
	return host, port
}

type API interface {
	Serve()
}

type api struct {
	host   string
	port   int
	server *echo.Echo
}

// @title DIT Backend API
// @version 1.0
// @description DIT Backend template for new backend projects
// @contact.name DIT - IFAL
// @contact.email wmrn1@aluno.ifal.edu.br
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewAPI(host string, port int) API {
	server := echo.New()
	server.Renderer = views.NewRenderer("src/apps/api/views/templates")
	server.HTTPErrorHandler = customHTTPErrorHandler
	return &api{host, port, server}
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	// Check if it's an API request (starts with /api)
	if strings.HasPrefix(c.Request().URL.Path, "/api") {
		c.Echo().DefaultHTTPErrorHandler(err, c)
		return
	}

	// For view requests, render custom error pages
	errorPage := "errors/500.html"
	title := "Erro Interno"

	switch code {
	case http.StatusNotFound:
		errorPage = "errors/404.html"
		title = "Página não encontrada"
	case http.StatusForbidden:
		errorPage = "errors/403.html"
		title = "Acesso Negado"
	case http.StatusUnauthorized:
		errorPage = "errors/401.html"
		title = "Não Autenticado"
	}

	// Attempt to get user info from context if available
	var userInfo interface{}
	if c.Get("user") != nil {
		userInfo = c.Get("user")
	}

	data := map[string]interface{}{
		"Title": title,
		"User":  userInfo,
	}

	if renderErr := c.Render(code, errorPage, data); renderErr != nil {
		c.Echo().DefaultHTTPErrorHandler(err, c)
	}
}

func (a *api) Serve() {
	a.setupStaticFiles()
	a.setupMiddlewares()
	a.loadRoutes()
	a.start()
}

func (a *api) setupStaticFiles() {
	a.server.Static("/static", "src/apps/api/static")
}

func (a *api) setupMiddlewares() {
	a.server.Use(middleware.Recover())
	a.server.Use(middlewares.LoggerMiddleware())
	a.server.Use(middlewares.CORSMiddleware())
}

func setUpAuth(cfg *config.AuthorizationConfig) {
	middlewares.AuthModelPath = cfg.AuthPaths.AuthModelPath
	middlewares.AuthPolicyPath = cfg.AuthPaths.AuthPolicyPath
}

func (a *api) rootEndpoint() *echo.Group {
	g := a.server.Group("/api")
	g.Use(middlewares.Authorize)
	return g
}

func (a *api) loadRoutes() {
	manager := routes.New()
	// API Routes
	manager.Load(a.rootEndpoint())

	// Frontend Views (at root /)
	routes.NewViewRouter().Load(a.server.Group(""))
}

func (a *api) start() {
	address := fmt.Sprintf("%s:%d", a.host, a.port)
	if err := a.server.Start(address); err != nil {
		a.server.Logger.Fatal(err)
	}
}
