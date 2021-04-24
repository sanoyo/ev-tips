package main

import (
	"log"
	"os"
	"strings"

	"github.com/mitchellh/cli"
)

type UpCommand struct {
}

func (c *UpCommand) Help() string {
	helpText := `
Usage: sql-migrate up [options] ...

  Migrates the database to the most recent version available.

Options:

  -config=dbconfig.yml   Configuration file to use.
  -env="development"     Environment.
  -limit=0               Limit the number of migrations (0 = unlimited).
  -dryrun                Don't apply migrations, just print them.

`
	return strings.TrimSpace(helpText)
}

func (c *UpCommand) Synopsis() string {
	return "Migrates the database to the most recent version available"
}

func (c *UpCommand) Run(args []string) int {
	return 0
}

func main() {
	c := cli.NewCLI("app", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"up": func() (cli.Command, error) {
			return &UpCommand{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
