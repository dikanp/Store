package routes

import (
	"AlterraStore/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.DELETE("/carts/:cartId/products/:productId", controllers.DeleteCartControllers)
	return e
}
