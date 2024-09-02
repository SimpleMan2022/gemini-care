package route

import "github.com/labstack/echo/v4"

func SetupRouter(e *echo.Echo) {
	auth := e.Group("/api/v1/auth")
	AuthRouter(auth)

	symptom := e.Group("/api/v1/symptom")
	SymptomRouter(symptom)
}
