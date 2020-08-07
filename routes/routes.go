package routes

import (
	"net/http"

	"github.com/at8109/go-echo-mysql/controllers"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, this is echo!")
	})

	e.GET("/employees", controllers.FetchAllEmployees)
	e.POST("/employees", controllers.StoreEmployee)
	e.PUT("/employees", controllers.UpdateEmployee)

	return e
}
