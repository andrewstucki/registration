package dump

import (
	"fmt"
	"io/ioutil"
	"os"

	"../../commands"
	"../../defaults"
	"../../utils"
)

var schemaPath string
var configPath string

var ConfigCmd = &commands.Command{
	Name: "dump:config",
	Usage: `
Dump:config dumps the default configuration file and schema file hard-coded
into the binary. The values are dumped into files specified by the -config
option and the -schema option respectively. Both can be modified
and used to customize the fields and look of the web form served by the
run:server command. Both a schema file and a config file must be in place
in order for the web server to run, so a dump:config command must be run
prior to the first time that the server is started.

Dump:config takes two flag:

  -schema
    The path to the schema output file (defaults to "schema.json").
  -config
    The path to the config output file (defaults to "config.json").
	`,
	Summary: "Dump default schema",
	Run:     dumpConfig,
}

func dumpConfig(cmd *commands.Command, args ...string) {
	err := ioutil.WriteFile(schemaPath, []byte(defaults.Schema), 0644)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(configPath, []byte(defaults.Config), 0644)
	if err != nil {
		panic(err)
	}
	utils.ChangeConsoleStyle("blue bold")
	fmt.Printf("Dumped default schema into '%s'!\n", schemaPath)
	fmt.Printf("Dumped default config into '%s'!\n", configPath)
	utils.ResetConsole()
	os.Exit(0)
}

func init() {
	ConfigCmd.Flag.StringVar(&schemaPath, "schema", "schema.json", "")
	ConfigCmd.Flag.StringVar(&configPath, "config", "config.json", "")
}
