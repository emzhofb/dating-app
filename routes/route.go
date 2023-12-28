package routes

import (
	"dating-app/constants"
	"dating-app/controllers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		message := `Mencintai diri sendiri berarti memahami bahwa kamu tidak perlu menjadi sempurna untuk menjadi baik.`
		return c.String(http.StatusOK, message)
	})

	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(constants.JWT_SECRET)))

	e.POST("/user/register", controllers.RegisterUserController)
	e.POST("/user/login", controllers.LoginUserController)

	eAuth.POST("/user/update", controllers.UpdateUserController)

	return e
}
