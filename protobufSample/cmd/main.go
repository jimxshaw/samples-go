package main

import (
	"flag"
	"fmt"
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
	return nil
}

func add(text string) error {
	task := &todo.Task{
		Text: text,
		Done: false,
	}

	fmt.Println(proto.MarshalTextString(task))

	return nil
}
