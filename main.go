package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

var todos []todo
var latestId = 0
var homeFolder string
var configFolder = homeFolder + ".config/todos-only/"

func help(args []string) {
	if len(args) <= 2 {
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

func getHome() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	homeFolder = dirname
}

func checkFolder() bool {
	if _, err := os.Stat(homeFolder + ".config"); os.IsNotExist(err) {
		if err := os.Mkdir(homeFolder+".config", os.ModePerm); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		return false
	}
	if _, err := os.Stat(homeFolder + ".config/todos-only"); os.IsNotExist(err) {
		if err := os.Mkdir(homeFolder+".config/todos-only", os.ModePerm); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		return false
	}
	return true
}

func save() {
	dataString := strconv.Itoa(latestId) + "\n"

	for _, todo := range todos {
		dataString += strconv.Itoa(todo.id) + "," + todo.todo + "," + strconv.FormatBool(todo.check) + "\n"
	}

	data := []byte(dataString)
	err := os.WriteFile(configFolder+"save.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	if !checkFolder() {
		checkFolder()
	}

	args := os.Args

	if len(args) == 1 {
		/* No Arguments */
		fmt.Printf("\nIt's cli todo program by hooss-only\n\n")
		help(args)
	} else {
		existsCommand := false
		cmd := strings.ToLower(args[1])
		var commandInfinitive command
		for _, c := range commands {
			if c.name == cmd {
				existsCommand = true
				commandInfinitive = c
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
			if len(args) < 3 {
				fmt.Println(commandInfinitive.usage)
				return
			}
			latestId++
			var todoText string
			for i, txt := range args[2:] {
				todoText += txt
				if !(i+1 == len(args[2:])) {
					todoText += " "
				}
			}

			todos = append(todos, todo{id: latestId, todo: todoText, check: false})

			fmt.Println(todos)
		}
	}
}
