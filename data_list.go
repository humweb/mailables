package mailables

// DataList structure for generating table components
type DataList struct {
	Title string
	Data  []DataListItem
}

func NewDataList(title string, data []DataListItem) DataList {
	return DataList{Title: title, Data: data}
}

type DataListItem struct {
	Key   string
	Value string
}
