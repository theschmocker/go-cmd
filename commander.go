package cmd

import (
	"flag"
	"fmt"
)

func NewCommander() Commander {
	return Commander{make(map[string]Command)}
}

type Commander struct {
	Commands map[string]Command
}

func (c *Commander) RegisterCommand(cmd Command) error {
	if _, ok := c.Commands[cmd.Name]; ok {
		return fmt.Errorf(`command "%s" already registered`, cmd.Name)
	}

	c.Commands[cmd.Name] = cmd
	return nil
}

func (c *Commander) Usage() {
	fmt.Print("Commands Available:\n\n")
	for _, cmd := range c.Commands {
		fmt.Printf("%s - %s\n", cmd.Name, cmd.Description)
	}
}

func (c *Commander) Execute(cmdName string, args []string) error {
	cmd, ok := c.Commands[cmdName]

	if !ok {
		return fmt.Errorf("command \"%s\" does not exist", cmdName)
	}

	var flags Flags
	cArgs := []string{}

	if len(args) > 2 {
		flags, cArgs = parseFlags(cmd.Flags, args[2:])
	}

	cmd.Execute(flags, cArgs)
	return nil
}

func parseFlags(fs []Flag, args []string) (Flags, []string) {
	flagSet := flag.NewFlagSet("flags", flag.ContinueOnError)
	strFlags := make(map[string]*string, 0)
	boolFlags := make(map[string]*bool, 0)

	for _, f := range fs {
		if f.IsBoolean {
			b := flagSet.Bool(f.Name, false, f.Description)
			flagSet.BoolVar(b, f.ShortName, false, f.Description)
			boolFlags[f.Name] = b
		} else {
			s := flagSet.String(f.Name, f.DefaultValue, f.Description)
			flagSet.StringVar(s, f.ShortName, f.DefaultValue, f.Description)
			strFlags[f.Name] = s
		}
	}

	flagSet.Parse(args)

	flags := Flags{
		String:  make(map[string]string),
		Boolean: make(map[string]bool),
	}

	for name, str := range strFlags {
		flags.String[name] = *str
	}

	for name, b := range boolFlags {
		flags.Boolean[name] = *b
	}

	return flags, flagSet.Args()
}
