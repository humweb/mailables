package mailables

// Basic imports
import (
	"github.com/stretchr/testify/suite"
	"html/template"
	"testing"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type MailableTestSuite struct {
	suite.Suite
	App      *Application
	Mailable *Mailable
	Body     *Body
}

func (suite *MailableTestSuite) SetupTest() {
	suite.App = NewApplication(
		"Test App",
		"http://test.com",
		"copyright test inc.",
		"",
	)
	suite.Mailable = NewMailable(
		suite.App,
		WithoutCSSInlining(),
	)

	suite.Body = NewBody()
}

func (suite *MailableTestSuite) TestApplicationProps() {
	suite.Equal("Test App", suite.App.Name)
	suite.Equal("http://test.com", suite.App.Link)
	suite.Equal("copyright test inc.", suite.App.Copyright)
	suite.Equal("", suite.App.Logo)
}

func (suite *MailableTestSuite) TestMailableProps() {
	suite.Equal(suite.App, suite.Mailable.App)
	suite.False(suite.Mailable.CSSInlining)
}

func (suite *MailableTestSuite) TestMailableThemeOption() {

	mailable := NewMailable(
		suite.App,
		WithoutCSSInlining(),
		WithTheme(&themeDef{}),
	)
	suite.Equal(mailable.Theme.Name(), "Test Theme")
	suite.Equal(mailable.Theme.Html(), "<p>Txt</p>")
	suite.Equal(mailable.Theme.Text(), "Txt")
}

func (suite *MailableTestSuite) TestBodySetters() {

	suite.Body.SetIntro("Intro content", ContentText)
	suite.Body.SetGreeting("Hey boo")
	suite.Body.SetName("boo")
	suite.Body.SetOutro("Outro Text", ContentText)
	suite.Body.SetSignature("Peace out!")
	suite.Body.SetTitle("Email body title")

	suite.Equal(suite.Body.Intro.Content, "Intro content")
	suite.Equal(suite.Body.Greeting, "Hey boo")
	suite.Equal(suite.Body.Name, "boo")
	suite.Equal(suite.Body.Outro.Content, "Outro Text")
	suite.Equal(suite.Body.Signature, "Peace out!")
	suite.Equal(suite.Body.Title, "Email body title")
}

func (suite *MailableTestSuite) TestSetTable() {
	table := NewTable("Test data table",
		[]TableRow{
			{{Key: "ID", Value: "1"}, {Key: "Title", Value: "foo"}},
			{{Key: "ID", Value: "2"}, {Key: "Title", Value: "bar"}},
		})
	suite.Body.AddTables(table)

	suite.Equal(suite.Body.Tables[0].Title, "Test data table")
	suite.Equal(suite.Body.Tables[0].Data[0][0].Key, "ID")
	suite.Equal(suite.Body.Tables[0].Data[0][0].Value, "1")
}

func (suite *MailableTestSuite) TestSetPanels() {
	panel1 := NewPanel("Title", NewContent(ContentText, "Body 1"), "50%")
	panel2 := NewPanel("", NewContent(ContentHTML, "Body 2"), "50%")
	suite.Body.AddPanels(panel1, panel2)

	suite.Equal(suite.Body.Panels[0].Title, "Title")
	suite.Equal(suite.Body.Panels[1].Title, "")
	suite.Equal(suite.Body.Panels[0].Body.Content, "Body 1")
	suite.Equal(suite.Body.Panels[1].Body.Content, "Body 2")
}

func (suite *MailableTestSuite) TestActions() {
	action1 := NewAction(Button{
		Text: "Help",
		Link: "http://test.com/help",
	}, "Get help")
	action2 := NewAction(Button{
		Text: "Verify Email",
		Link: "http://test.com/verify/123",
	}, "Click here to verify email")

	suite.Body.AddActions(action1, action2)

	suite.Equal(suite.Body.Actions[0].Button.Text, "Help")
	suite.Equal(suite.Body.Actions[0].Button.Link, "http://test.com/help")
	suite.Equal(suite.Body.Actions[1].Button.Text, "Verify Email")
	suite.Equal(suite.Body.Actions[1].Instructions, "Click here to verify email")
}

func (suite *MailableTestSuite) TestContentMarkdown() {
	content := NewContent(ContentMarkdown, "**Bolded**")
	content2 := NewContent(ContentHTML, "<b>Bolded</b>")

	suite.Equal(template.HTML("<p><strong>Bolded</strong></p>\n"), content.ToHTML())
	suite.Equal(template.HTML("<b>Bolded</b>"), content2.ToHTML())
}

func (suite *MailableTestSuite) TestDefaultTemplate() {
	tmpl := ThemeDefault{}

	suite.IsType("string", tmpl.Html())
	suite.IsType("string", tmpl.Text())
	suite.Equal("default", tmpl.Name())
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestMailableTestSuite(t *testing.T) {
	suite.Run(t, new(MailableTestSuite))
}

type themeDef struct{}

func (t *themeDef) Name() string {
	return "Test Theme"
}
func (t *themeDef) Html() string {
	return "<p>Txt</p>"
}
func (t *themeDef) Text() string {
	return "Txt"
}
