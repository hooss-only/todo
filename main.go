package main

import (
	"container/list"
	"fmt"
	"os"
	"strings"
)

type command struct {
	name        string
	description string
	usage       string
}

type todo struct {
	id    int
	todo  string
	check bool
}

var commands = [...]command{
	{name: "help", description: "give helps about commands", usage: "todo help [command]"},
	{name: "add", description: "add new todo", usage: "todo add <to-do>"},
}

var todos = list.New()

func help(args []string) {
	if len(args) == 2 {
		fmt.Printf("Usage:\n\ttodo <command> [args..]\n\n")

		fmt.Println("Commands:")
		for _, c := range commands {
			fmt.Println("\t"+c.name, "|", c.description)
		}
		fmt.Println("")
	} else {
		findingCommand := args[2]
		var cmd command

		existsCommand := false

		for _, c := range commands {
			if c.name == findingCommand {
				cmd = c
				existsCommand = true
			}
		}

		if !existsCommand {
			fmt.Printf("%s: unknown command.\n", findingCommand)
			return
		}

		fmt.Printf("\n%s\n\t%s\n\t%s\n\n", cmd.name, cmd.description, cmd.usage)
	}

	return
}

func main() {
	args := os.Args

	if len(args) == 1 {
		/* No Arguments */
		fmt.Printf("\nIt's cli todo program by hooss-only\n\n")
		help(args)
	} else {
		existsCommand := false
		cmd := strings.ToLower(args[1])
		for _, c := range commands {
			if c.name == cmd {
				existsCommand = true
			}
		}

		if !existsCommand {
			fmt.Println(cmd + ": unknown command")
			return
		}

		switch cmd {
		case "help":
			help(args)
		case "add":

		}
	}
}
