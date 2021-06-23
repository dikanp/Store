package main

import "github.com/labstack/echo"

func main() {
	e := echo.New()

	e.GET("/", DeleteProductController)
	e.Start(":8080")
}