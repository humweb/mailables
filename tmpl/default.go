package tmpl

import (
	"embed"
	"html/template"
)

//go:embed *.gohtml
var templates embed.FS

type ThemeDefault struct {
	parsedTemplates *template.Template
}

var templateFuncs = template.FuncMap{
	"url": func(s string) template.URL {
		return template.URL(s)
	},
}

func (t *ThemeDefault) Name() string {
	return "default"
}

func (t *ThemeDefault) Text() string {
	return ""
}

func (t *ThemeDefault) Html() *template.Template {
	if t.parsedTemplates == nil {
		t.parsedTemplates = template.Must(template.New("index.gohtml").
			Funcs(templateFuncs).
			ParseFS(templates, "css.gohtml", "actions.gohtml", "data_list.gohtml", "panels.gohtml", "tables.gohtml", "index.gohtml"))
	}
	return t.parsedTemplates
}
