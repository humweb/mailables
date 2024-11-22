package tmpl

import "html/template"

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
		t.parsedTemplates = template.Must(template.New(t.Name()).
			Funcs(templateFuncs).
			ParseFiles("css.gohtml", "actions.gohtml", "data_list.gohtml", "panels.gohtml", "tables.gohtml", "index.gohtml"))
	}
	return t.parsedTemplates
}
