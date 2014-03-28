package commands

import (
	"flag"
	"fmt"
	"os"
)

type Command struct {
	Run  func(cmd *Command, args ...string)
	Flag flag.FlagSet

	Name  string
	Usage string

	Summary string
}

func (c *Command) Exec(args []string) {
	c.Flag.Usage = func() {
		fmt.Print(c.Usage)
		os.Exit(1)
	}
	c.Flag.Parse(args)
	c.Run(c, c.Flag.Args()...)
}
