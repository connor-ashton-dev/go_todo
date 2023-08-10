package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/connor-ashton-dev/todo_cli"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "Add a new todo item")
	completed := flag.Int("complete", 0, "Mark a todo item as completed")
	del := flag.Int("del", 0, "Delete a todo item")
	list := flag.Bool("list", false, "List all todo items")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Println(err)
		return
	}

	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Println(err)
			return
		}
		todos.Add(task)
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Println(err)
			return
		}

	case *completed > 0:
		err := todos.Complete(*completed)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Println(err)
			return
		}

	case *del > 0:
		err := todos.Delete(*del)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Println(err)
			return
		}
	case *list:
		todos.Print()

	default:
		fmt.Println("No command specified")
		return
	}
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("No input provided")
	}
	return text, nil
}
