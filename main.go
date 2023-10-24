package main

import (
	"fmt"
	"os"
	"strings"
)

type command struct {
	name        string
	description string
}

func main() {
	commands := [...]command{{name: "help", description: "give helps about commands"}, {name: "add", description: "add new todo"}}
	args := os.Args

	if len(args) == 1 {
		/* No Arguments */
		fmt.Printf("It's cli todo program by hooss-only\n\n")
		fmt.Printf("Usage:\n\ttodo <command> [args..]\n\n")

		fmt.Println("Commands:")
		for _, cmd := range commands {
			fmt.Println("\t"+cmd.name, "|", cmd.description)
		}
		fmt.Println("")
	} else {
		existsCommand := false
		command := strings.ToLower(args[1])

		for _, cmd := range commands {
			if cmd.name == command {
				existsCommand = true
			}
		}

		if !existsCommand {
			fmt.Println(command + ": unknown command")
			return
		}
	}
}
