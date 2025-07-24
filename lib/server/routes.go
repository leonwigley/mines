package server

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	render "mines/app/views"
	"mines/config"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Serve static assets
	// fileServer := http.FileServer(http.Dir("cmd/web/assets"))
	// e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", fileServer)))
	e.Static("/assets", "public/")

	// Routes
	e.GET("/", func(c echo.Context) error {
		return render.Index(c.Response().Writer, map[string]interface{}{
			"AppName": config.GlobalConfig.AppName,
			"Title":   "Modern web apps, the simple way.",
		}, "")
	})

	e.GET("/profile", func(c echo.Context) error {
		return render.About(c.Response().Writer, map[string]interface{}{
			"AppName": config.GlobalConfig.AppName,
			"Title":   "About us",
			"Name":    "Leon",
		}, "")
	})

	e.GET("/health", s.healthHandler)

	return e
}

// Renderer
type TemplateRenderer struct {
	Templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if err := t.Templates.ExecuteTemplate(w, name, data); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

// Health function
func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
