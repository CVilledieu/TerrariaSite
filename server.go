package main

import (
	//info "Terraria/info"

	"io"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)

}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("site/*.html")),
	}
}

type Page struct {
	Search  Search
	Results Results
}

func newPage() Page {
	return Page{
		//Search: searchResults("set up"),
	}
}

type Search struct {
	lookup  string
	catalog string
}

type Results struct {
	Id      int64
	Name    string
	Cost    int64
	ToolTip string
}

func startServer() {
	e := echo.New()
	e.Use(middleware.Logger())
	page := newPage()

	e.Renderer = newTemplate()
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})

	/*
		e.GET("/search", func(c echo.Context) error {
			results := searchResults("name")
			_, err := fmt.Println(results)
			if err != nil {
				log.Fatal(err)
			}
			//return c.Render(200, "results", results)
		})

		e.POST("/count", func(c echo.Context) error {
			//Count.Count++
			//return c.Render(200, "count", Count)
		})

		e.Logger.Fatal(e.Start(":8080"))
	*/
}
