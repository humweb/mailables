package mailables

import (
	"bytes"
	"github.com/vanng822/go-premailer/premailer"
	"html/template"
)

// Mailable
type Mailable struct {
	App            *Application
	Theme          Theme
	CSSInlining    bool
	CachedTemplate *template.Template
}
type Page struct {
	Application *Application
	Body        *Body
}

var templateFuncs = template.FuncMap{
	"url": func(s string) template.URL {
		return template.URL(s)
	},
}

func (m *Mailable) ToHTML(body *Body) (string, error) {
	var err error
	// Parse and cache template
	//if m.CachedTemplate == nil {
	m.CachedTemplate, err = template.New("mailable").Funcs(templateFuncs).Parse(m.Theme.Html())
	if err != nil {
		return "", err
	}
	//}
	var b bytes.Buffer
	if err := m.CachedTemplate.Execute(&b, &Page{
		Application: m.App,
		Body:        body,
	}); err != nil {
		return "", err
	}

	res := b.String()
	if m.CSSInlining {
		if prem, err := premailer.NewPremailerFromString(res, premailer.NewOptions()); err == nil {
			return prem.Transform()
		}
	}

	return res, nil
}

// MailableOption function types to set options for mailable
type MailableOption func(*Mailable)

// NewMailable init new mailable
func NewMailable(app *Application, opts ...MailableOption) *Mailable {
	m := &Mailable{
		App:         app,
		CSSInlining: true,
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// WithTheme option allows you to override the default theme
func WithTheme(theme Theme) MailableOption {
	return func(m *Mailable) {
		m.Theme = theme
	}
}

// WithoutCSSInlining option disables CSS inlining when parsing the theme
func WithoutCSSInlining() MailableOption {
	return func(m *Mailable) {
		m.CSSInlining = false
	}
}

// Application holds global values used for building templates
type Application struct {
	Name      string
	Link      string
	Logo      string
	Copyright string
}

// NewApplication init new Application
func NewApplication(name, link, copyright, logo string) *Application {
	return &Application{
		Name:      name,
		Link:      link,
		Logo:      logo,
		Copyright: copyright,
	}
}

// Theme contract for theme expansion
type Theme interface {
	Name() string
	Html() string
	Text() string
}

// Body this holds the main data for the mailable. It is used to build each email section
type Body struct {
	Title     string // Title replaces the greeting+name when set
	Name      string // The name of the contacted person
	Greeting  string // Greeting for the contacted person (default to 'Hi')
	Intro     Content
	Panels    []Panel
	Tables    []Table
	Actions   []Action // Actions are a list of actions that the user will be able to execute via a button click
	Outro     Content
	Signature string // Signature for the contacted person (default to 'Yours truly')
}

// NewBody init new body struct
func NewBody() *Body {
	return &Body{}
}

// SetName set name for body
func (b *Body) SetName(name string) *Body {
	b.Name = name
	return b
}

// SetSignature set signature for body
func (b *Body) SetSignature(signature string) *Body {
	b.Signature = signature
	return b
}

// SetTitle set title for body
func (b *Body) SetTitle(title string) *Body {
	b.Title = title
	return b
}

// SetGreeting ste greeting for body
func (b *Body) SetGreeting(greeting string) *Body {
	b.Greeting = greeting
	return b
}

// SetIntro set intro for body
func (b *Body) SetIntro(intro string, format ContentFormat) *Body {
	b.Intro = Content{
		Format:  format,
		Content: intro,
	}
	return b
}

// SetOutro set outro for body
func (b *Body) SetOutro(outro string, format ContentFormat) *Body {
	b.Outro = Content{
		Format:  format,
		Content: outro,
	}
	return b
}

// AddPanels set panel(s) for body
func (b *Body) AddPanels(panels ...Panel) *Body {
	b.Panels = append(b.Panels, panels...)
	return b
}

// AddTables et tables(s) for body
func (b *Body) AddTables(tables ...Table) *Body {
	b.Tables = append(b.Tables,
		tables...)
	return b
}

// AddActions set action(s) for body
func (b *Body) AddActions(actions ...Action) *Body {
	b.Actions = append(b.Actions, actions...)
	return b
}

//t, err := template.New("hermes").Funcs(sprig.FuncMap()).Funcs(templateFuncs).Funcs(template.FuncMap{
//"safe": func(s string) template.HTML { return template.HTML(s) }, // Used for keeping comments in generated template
//}).Parse(tplt)
