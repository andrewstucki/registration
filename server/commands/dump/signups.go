package dump

import (
	"fmt"
	"os"

	"../../commands"
	"../../database"
	"../../utils"
)

var csvPath string
var databasePath string

var SignupsCmd = &commands.Command{
	Name: "dump:signups",
	Usage: `
Dump:signups dumps all of the signups in a given database into a csv file
for easy exporting.

Two flags are available:

  -csv
    The path to the csv output file (defaults to "signups.csv").
  -db
    The path to the registration database (defaults to "registration.db").
  `,
	Summary: "Dump all signups",
	Run:     dumpSignups,
}

func dumpSignups(cmd *commands.Command, args ...string) {
	err := database.Initialize(databasePath)
	if err != nil {
		panic(err)
	}
	err = database.DumpSignups(csvPath)
	if err != nil {
		panic(err)
	}
	utils.ChangeConsoleStyle("blue bold")
	fmt.Printf("Dumped signups to '%s'!\n", csvPath)
	utils.ResetConsole()
	os.Exit(0)
}

func init() {
	SignupsCmd.Flag.StringVar(&csvPath, "csv", "signups.csv", "")
	SignupsCmd.Flag.StringVar(&databasePath, "db", "registration.db", "")
}
