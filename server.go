package main

import (
	//info "Terraria/info"
	"Terraria/info"
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Templates struct {
	templates *template.Template
}

type (
	PlayerItems    info.Player
	NonPlayerItems info.NonPlayer
	NPC            info.NPC
)

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
		Search: searchResults("set up"),
	}
}

type Search struct {
	lookup  string
	catalog string
}

func searchResults(name string) Search {
	if name == "set up" {
		return Search{
			lookup:  "",
			catalog: "",
		}
	}
	results := //search item
	return results	
}

type Results struct {
	Id      int64
	Name    string
	Cost    int64
	ToolTip string
}


