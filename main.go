package main

import "fmt"
import "os"

type command struct {
	name				string
	description string
}

func main() {
	commands := [...]command{{name: "help", description: "give helps about commands"}, {name: "add", description: "add new todo"}}
	args := os.Args

	if len(args) == 1 {
		/* No Arguments */
		fmt.Println("It's cli todo program by hooss-only\n");
		fmt.Println("Usage:\n\ttodo <command> [args..]\n");
	
		fmt.Println("Commands:")
		for _, cmd := range commands {
			fmt.Println("\t"+cmd.name, "|", cmd.description);
		}
		fmt.Println("")
	} else {
		commandNames := [...]string{}


	}
}
