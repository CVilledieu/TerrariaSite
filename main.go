package main

import (
	info "Terraria/info"

	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	item, err := info.GetToolInfo("Tin Pickaxe")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(item)

	e := echo.New()
	e.Use(middleware.Logger())
	page := newPage()

	e.Renderer = newTemplate()
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})

	e.GET("/tools", func(c echo.Context) error {
		return c.Render(200, "tools", Count)
	})

	e.POST("/count", func(c echo.Context) error {
		Count.Count++
		return c.Render(200, "count", Count)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
