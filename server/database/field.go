package database

type field struct {
	Type       string
	Required   bool
	Group      string
	Summarize  bool
	Values     []string
	FormType   string
	Message    string
	Validation string
	Prepend    string
	ExtraClass string
	Label      string
}
