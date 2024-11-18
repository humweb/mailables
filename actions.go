package mailables

// Action
type Action struct {
	Instructions string
	Button       Button
}

// NewAction
func NewAction(button Button, instructions string) Action {
	return Action{Instructions: instructions, Button: button}
}

// Button defines a action Button
type Button struct {
	Variant string
	Text    string
	Link    string
}
