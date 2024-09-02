package main

import (
	"gemini-care/api/route"
	"gemini-care/bootstrap"
	"github.com/labstack/echo/v4"
)

func main() {
	bootstrap.Initialize()
	e := echo.New()
	route.SetupRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}
