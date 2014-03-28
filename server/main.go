package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"./commands"
	"./commands/dump"
	"./commands/run"
	"./globals"
)

var Commands = []*commands.Command{
	dump.ConfigCmd,
	dump.SignupsCmd,
	run.ServerCmd,
}

var usagePrefix = `
register is a portable registration server.

Usage:
  ./register [options] <subcommand> [subcommand options]

Options:
`
var usageTmpl = template.Must(template.New("usage").Parse(`
Commands:{{range .}}
  {{.Name | printf "%-14s"}} {{.Summary}}{{end}}
`))

func usage() {
	fmt.Print(usagePrefix)
	flag.PrintDefaults()
	usageTmpl.Execute(os.Stdout, Commands)
}

func main() {
	flag.BoolVar(&globals.Debug, "debug", false, "Turn debugging on.")
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 || args[0] == "-h" {
		flag.Usage()
		return
	}

	var cmd *commands.Command
	name := args[0]
	for _, c := range Commands {
		if strings.HasPrefix(c.Name, name) {
			cmd = c
			break
		}
	}

	if cmd == nil {
		fmt.Printf("error: unknown command %q\n", name)
		flag.Usage()
		os.Exit(1)
	}

	cmd.Exec(args[1:])
}
