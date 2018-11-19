package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ymr-39/monkey-lang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user)
	repl.Start(os.Stdin, os.Stdout)
}
