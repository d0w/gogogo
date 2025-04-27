package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is a random programming language\n",
		user.Username)

	fmt.Printf("Type in some commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
