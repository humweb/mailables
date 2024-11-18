package mailables

import (
	"github.com/russross/blackfriday/v2"
	"html/template"
)

// ContentFormat content format type
type ContentFormat string

const ContentMarkdown ContentFormat = "markdown"
const ContentHTML ContentFormat = "html"
const ContentText ContentFormat = "text"

// Content generic content structure to describe content and what format the content is
type Content struct {
	Format  ContentFormat
	Content string
}

func (c *Content) ToHTML() template.HTML {
	if c.Format == ContentMarkdown {
		return template.HTML(blackfriday.Run([]byte(c.Content)))
	}
	return template.HTML(c.Content)
}

func NewContent(format ContentFormat, content string) *Content {
	return &Content{Format: format, Content: content}
}

// Panel simple title and body panel to display panel/card data
type Panel struct {
	Title string
	Body  *Content
}

func NewPanel(title string, body *Content) Panel {
	return Panel{Title: title, Body: body}
}
