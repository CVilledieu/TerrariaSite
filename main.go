package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	page := newPage()

	e.Renderer = newTemplate()
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})

	e.GET("/search", func(c echo.Context) error {
		results := searchResults("name")
		return c.Render(200, "results", results)
	})

	e.POST("/count", func(c echo.Context) error {
		Count.Count++
		return c.Render(200, "count", Count)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
