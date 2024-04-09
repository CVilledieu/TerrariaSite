package main

import (
	info "Terraria/info"
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

type Page struct {
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)

}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("site/*.html")),
	}
}

func newSearch(itemType, itemName string) (PlayerItems, error) {
	if itemType == "tool" {
		newSearch := PlayerItems{}
		item, err := info.GetToolInfo(itemName)
		if err != nil {
			return PlayerItems{}, err
		}
		newSearch.ToolData = item
		return newSearch, nil
	}

	return item, nil
}

func newPage() Page {
	return Page{}
}
