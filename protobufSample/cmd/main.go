package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/jimxshaw/samples-go/protobufSample/todo"
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}

	var err error

	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list()
	case "add":
		err = add(strings.Join(flag.Args()[1:], " "))
	default:
		err = fmt.Errorf("unknown subcommand %s", cmd)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

const dbPath = "mydb.pb"

func list() error {
	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		return fmt.Errorf("Could not read file %s: %v", dbPath, err)
	}

	for {
		var task todo.Task

		if err := proto.Unmarshal(b, &task); err == io.EOF {
			return nil
		} else if err != nil {
			return fmt.Errorf("Could not read task: %v", err)
		}

		if task.Done {
			fmt.Printf("[O]")
		} else {
			fmt.Printf("[X]")
		}

		fmt.Printf(" %s\n", task.Text)
	}

}

func add(text string) error {
	task := &todo.Task{
		Text: text,
		Done: false,
	}

	b, err := proto.Marshal(task)
	if err != nil {
		return fmt.Errorf("Could not encode task: %v", err)
	}

	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("Could not open %s: %v", dbPath, err)
	}

	_, err = f.Write(b)
	if err != nil {
		return fmt.Errorf("Could not write task to file: %v", err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("Could not close file %s: %v", dbPath, err)
	}

	return nil
}
