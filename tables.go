package mailables

// Table structure for generating table components
type Table struct {
	Title string
	Data  []TableRow
}

func NewTable(title string, data []TableRow) Table {
	return Table{Title: title, Data: data}
}

type TableColumn struct {
	Key   string
	Value string
}

// TableRow structure for generating a table row
type TableRow []TableColumn
