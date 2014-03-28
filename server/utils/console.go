package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func ChangeConsoleStyle(options string) {
	fmt.Printf("%s", styleText(options, ""))
}

func ResetConsole() {
	fmt.Printf("%s", styleText("reset", ""))
}

func styleText(options, text string) (result string) {
	styles := regexp.MustCompile("[^\\w]+").Split(options, -1)
	result = text

	for _, style := range styles {
		result = fmt.Sprintf("%s%s", ansiCodes[strings.Title(style)], result)
	}

	return result
}

var ansiCodes = map[string]string{
	"Reset":            "\033[0m",
	"Bold":             "\033[1m",
	"Italic":           "\033[3m",
	"Blink":            "\033[5m",
	"Underline":        "\033[4m",
	"UnderlineOff":     "\033[24m",
	"Inverse":          "\033[7m",
	"InverseOff":       "\033[27m",
	"Strikethrough":    "\033[9m",
	"StrikethroughOff": "\033[29m",

	"Def":     "\033[39m",
	"White":   "\033[37m",
	"Black":   "\033[30m",
	"Grey":    "\x1B[90m",
	"Red":     "\033[31m",
	"Green":   "\033[32m",
	"Blue":    "\033[34m",
	"Yellow":  "\033[33m",
	"Magenta": "\033[35m",
	"Cyan":    "\033[36m",

	"DefBg":     "\033[49m",
	"WhiteBg":   "\033[47m",
	"BlackBg":   "\033[40m",
	"RedBg":     "\033[41m",
	"GreenBg":   "\033[42m",
	"BlueBg":    "\033[44m",
	"YellowBg":  "\033[43m",
	"MagentaBg": "\033[45m",
	"CyanBg":    "\033[46m",
}

type Table struct {
	Fields     []string
	Rows       []map[string]string
	fieldSizes map[string]int
}

func NewTable(fields []string) *Table {
	return &Table{
		Fields:     fields,
		Rows:       make([]map[string]string, 0),
		fieldSizes: make(map[string]int),
	}
}

func (t *Table) AddRow(row map[string]interface{}) {
	newRow := make(map[string]string)
	for _, k := range t.Fields {
		v := row[k]
		// If is not nil format
		// else value is empty string
		var val string
		if v == nil {
			val = ""
		} else {
			val = fmt.Sprintf("%v", v)
		}

		valLen := len(val)
		// align to field name length
		if valLen < len(k) {
			valLen = len(k)
		}
		valLen += 2 // + 2 spaces
		if t.fieldSizes[k] < valLen {
			t.fieldSizes[k] = valLen
		}
		newRow[k] = val
	}
	if len(newRow) > 0 {
		t.Rows = append(t.Rows, newRow)
	}
}

func (t *Table) Print() {
	if len(t.Rows) == 0 {
		return
	}
	t.printDash()
	fmt.Println(t.getHead())
	t.printDash()
	for _, r := range t.Rows {
		fmt.Println(t.rowString(r))
		t.printDash()
	}
}

func (t *Table) getHead() string {
	s := "|"
	for _, name := range t.Fields {
		s += t.fieldString(name, strings.Title(name)) + "|"
	}
	return s
}

func (t *Table) rowString(row map[string]string) string {
	s := "|"
	for _, name := range t.Fields {
		value := row[name]
		s += t.fieldString(name, value) + "|"
	}
	return s
}

func (t *Table) fieldString(name, value string) string {
	value = fmt.Sprintf(" %s ", value)
	spacesLeft := t.fieldSizes[name] - len(value)
	if spacesLeft > 0 {
		for i := 0; i < spacesLeft; i++ {
			value += " "
		}
	}
	return value
}

func (t *Table) printDash() {
	s := "|"
	for i := 0; i < t.lineLength()-2; i++ {
		s += "-"
	}
	s += "|"
	fmt.Println(s)
}

func (t *Table) lineLength() (sum int) {
	sum = 1
	for _, l := range t.fieldSizes {
		sum += l + 1
	}
	return sum
}

func PrintRow(fields []string, row map[string]interface{}) {
	table := NewTable(fields)
	// add row
	table.AddRow(row)
	// And display table
	table.Print()
}

func PrintTable(fields []string, rows []map[string]interface{}) {
	// Construct a table
	table := NewTable(fields)
	for _, r := range rows {
		table.AddRow(r)
	}

	// And display table
	table.Print()
}
