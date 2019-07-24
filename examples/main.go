package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/theschmocker/cmd"
)

func main() {
	cmder := cmd.NewCommander()

	cat := cmd.Command{
		Name:        "cat",
		Description: "prints main.go to the console",
		Flags: []cmd.Flag{{
			Name:         "f",
			DefaultValue: "",
			Description:  "file name",
		}},
		Execute: Cat,
	}

	greet := cmd.Command{
		Name:        "greet",
		Description: "Outputs the first positional argument",
		Flags: []cmd.Flag{{
			Name:        "l",
			Description: "makes it loud",
			IsBoolean:   true,
		}},
		Execute: Greet,
	}

	//cmder.RegisterCommand(version)
	cmder.RegisterCommand(cat)
	cmder.RegisterCommand(greet)

	cmdName := os.Args[1]
	cmder.Execute(cmdName, os.Args)
}

func Greet(flags cmd.Flags, args []string) {
	loud := flags.Boolean["l"]
	greeting := args[0]

	if greeting == "" {
		fmt.Println("Pass a greeting!")
		os.Exit(1)
	}

	if loud {
		greeting = strings.ToUpper(greeting)
	}

	fmt.Println(greeting)
}

func Cat(flags cmd.Flags, args []string) {
	fileName := flags.String["f"]

	if fileName == "" {
		log.Fatal("blaaaaaaaah")
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func Version() {
	fmt.Println("version 1.0")
}
