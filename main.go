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
	{name: "list", description: "show todo list", usage: "todo list"},
	{name: "check", description: "check or uncheck a todo", usage: "todo check <id>"},
	{name: "del", description: "delete todo", usage: "todo del <id> [id] [id] ... || todo del checked"},
	{name: "reset", description: "reset todos and accumulated id amount", usage: "todo reset"},
}

var todos []todo
var latestId = 0
var homeFolder string
var configFolder string

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

func initFolderPath() {
	dirname, err := os.UserHomeDir()
	errCheck(err)
	homeFolder = dirname
	configFolder = homeFolder + "/.config/todos-only/"
}

func checkFolder() bool {
	if _, err := os.Stat(homeFolder + "/.config"); os.IsNotExist(err) {
		if err := os.Mkdir(homeFolder+"/.config", os.ModePerm); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		return false
	}
	if _, err := os.Stat(homeFolder + "/.config/todos-only"); os.IsNotExist(err) {
		if err := os.Mkdir(homeFolder+"/.config/todos-only", os.ModePerm); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		return false
	}
	if _, err := os.Stat(configFolder + "save.txt"); os.IsNotExist(err) {
		file, err := os.Create(configFolder + "save.txt")
		errCheck(err)
		defer file.Close()

		data := []byte("0")
		err = os.WriteFile(configFolder+"save.txt", data, 0644)
		errCheck(err)
		return false
	}
	return true
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func save() {
	dataString := strconv.Itoa(latestId) + "\n"

	for _, todo := range todos {
		dataString += strconv.Itoa(todo.id) + "," + todo.todo + "," + strconv.FormatBool(todo.check) + "\n"
	}

	data := []byte(dataString)
	err := os.WriteFile(configFolder+"save.txt", data, 0644)
	errCheck(err)
}

func load() {
	data, err := os.ReadFile(configFolder + "save.txt")
	errCheck(err)
	dataString := string(data)

	for idx, context := range strings.Split(dataString, "\n") {
		if idx == 0 {
			latestId, err = strconv.Atoi(context)
			errCheck(err)
		} else {
			todoSplit := strings.Split(context, ",")

			if len(todoSplit) == 1 {
				return
			}

			id, err := strconv.Atoi(todoSplit[0])
			errCheck(err)

			todoText := todoSplit[1]

			check, err := strconv.ParseBool(todoSplit[2])
			errCheck(err)

			todos = append(todos, todo{id: id, todo: todoText, check: check})
		}
	}
}

func getIndexById(id int) int {
	for idx, ctx := range todos {
		if ctx.id == id {
			return idx
		}
	}
	return -1
}

func remove(s []todo, i int) []todo {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	initFolderPath()

	if !checkFolder() {
		checkFolder()
	}

	load()

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
		case "list":
			if len(todos) == 0 {
				fmt.Println("its too clean")
			}
			for _, ctx := range todos {
				checkText := ""

				if ctx.check {
					checkText = "✓"
				} else {
					checkText = " "
				}

				fmt.Printf("[%d] [%s]      %s\n", ctx.id, checkText, ctx.todo)
			}
		case "check":
			if len(args) != 3 {
				fmt.Println(commands[3].usage)
				return
			} else {
				id, err := strconv.Atoi(args[2])
				errCheck(err)
				index := getIndexById(id)

				todos[index].check = !todos[index].check
			}
		case "del":
			if len(args) < 3 {
				fmt.Println(commands[4].usage)
				return
			} else {
				if args[2] == "checked" {
					var ids []int

					for _, ctx := range todos {
						if ctx.check {
							ids = append(ids, ctx.id)
						}
					}

					for _, id := range ids {
						todos = remove(todos, getIndexById(id))
					}
				} else {
					for _, ctx := range args[2:] {
						id, err := strconv.Atoi(ctx)
						errCheck(err)
						todos = remove(todos, getIndexById(id))
					}
				}
			}
		case "reset":
			latestId = 0
			todos = []todo{}
		}
		save()
	}
}
