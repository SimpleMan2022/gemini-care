package main

import (
	"gemini-care/api/route"
	"gemini-care/bootstrap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	bootstrap.Initialize()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	route.SetupRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}
