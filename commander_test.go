package cmd

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestRegisterCommand(t *testing.T) {
	cmder := NewCommander()

	c := Command{
		Name:        "test",
		Description: "A fake command for testing",
		Flags:       []Flag{},
		Execute:     func(flags Flags, args []string) {},
	}

	cmder.RegisterCommand(c)

	r := cmder.Commands["test"]

	if reflect.DeepEqual(c, r) {
		t.Errorf("Expeced command %v to equal %v", cmder.Commands["test"], c)
	}
}

func TestRegisterCommandWithSameNameReturnsError(t *testing.T) {
	cmder := NewCommander()

	c := Command{
		Name:        "test",
		Description: "A fake command for testing",
		Flags:       []Flag{},
		Execute:     func(flags Flags, args []string) {},
	}

	c2 := Command{
		Name:        "test",
		Description: "A duplicate fake command for testing",
		Flags:       []Flag{},
		Execute:     func(flags Flags, args []string) {},
	}

	cmder.RegisterCommand(c)
	err := cmder.RegisterCommand(c2)

	if err == nil {
		t.Errorf("Expected %s (%s) to conflict with %s (%s)", c.Name, c.Description, c2.Name, c2.Description)
	}
}

func TestExecutingNonExistentCommandReturnsError(t *testing.T) {
	cmder := NewCommander()
	err := cmder.Execute("fake", []string{})

	if err == nil {
		t.Error("Expected Execute to return an error")
	}
}

func TestExecutingCommandReturnsNil(t *testing.T) {

	cmder := NewCommander()

	c := Command{
		Name:        "test",
		Description: "A fake command for testing",
		Flags:       []Flag{},
		Execute:     func(flags Flags, args []string) {},
	}

	cmder.RegisterCommand(c)

	err := cmder.Execute("test", []string{})

	if err != nil {
		t.Error("Expected Execute to return nil")
	}
}

func ExampleExecutingCommandWithArgs() {
	cmder := NewCommander()

	c := Command{
		Name:        "greet",
		Description: "Greet someone (loudly?)",
		Flags: []Flag{
			{
				Name:        "yell",
				ShortName:   "y",
				Description: "greet loudly",
				IsBoolean:   true,
			},
		},
		Execute: func(flags Flags, args []string) {
			greeting := "Hello, " + args[0]
			yell := flags.Boolean["yell"]

			if yell {
				greeting = strings.ToUpper(greeting) + "!"
			}

			fmt.Println(greeting)
		},
	}

	cmder.RegisterCommand(c)

	args := []string{"cli", "greet", "-y", "April"}

	cmder.Execute(args[1], args)

	// Output:
	// HELLO, APRIL!
}

func TestParseFlagsWithStringFlags(t *testing.T) {
	f := []Flag{
		{
			Name:         "flag",
			ShortName:    "f",
			Description:  "A test string flag",
			DefaultValue: "test",
		},
		{
			Name:         "another-flag",
			ShortName:    "a",
			Description:  "Another test string flag",
			DefaultValue: "test 2",
		},
	}

	args := []string{"--flag", "flag value", "--another-flag", "another value"}

	flags, _ := parseFlags(f, args)

	if len(flags.String) != 2 {
		t.Errorf("Expected two string flags")
	}

	for i, flag := range f {
		val, ok := flags.String[flag.Name]
		if !ok {
			t.Errorf(`Expected there to a value at flags.String["%s"]`, flag.Name)
		} else if val != args[i*2+1] {
			t.Errorf(`Expected "%s" to be "%s"`, val, args[i*2+1])
		}
	}
}

func TestParseFlagsWithBooleanFlags(t *testing.T) {
	f := []Flag{
		{
			Name:        "is-cool",
			ShortName:   "c",
			Description: "A test string flag",
			IsBoolean:   true,
		},
	}

	args := []string{"--is-cool"}

	flags, _ := parseFlags(f, args)

	if len(flags.Boolean) != 1 {
		t.Error("Expected one boolean flags")
	}

	if !flags.Boolean["is-cool"] {
		t.Error("Expected is-cool to be true")
	}
}

func TestParseFlagsWithShortNames(t *testing.T) {
	f := []Flag{
		{
			Name:         "flag",
			ShortName:    "f",
			Description:  "A test string flag",
			DefaultValue: "test",
		},
	}

	args := []string{"-f", "flag value"}

	flags, _ := parseFlags(f, args)

	if val, ok := flags.String["flag"]; !ok || val != "flag value" {
		t.Error(`Expected short name for "flag" ("f") to be parsed correctly`)
	}
}

func ExampleUseage() {
	cmder := NewCommander()

	c := Command{
		Name:        "test",
		Description: "A fake command for testing",
		Flags:       []Flag{},
		Execute:     func(flags Flags, args []string) {},
	}

	cmder.RegisterCommand(c)

	cmder.Usage()
	// Output:
	// Commands Available:
	//
	// test - A fake command for testing
}
