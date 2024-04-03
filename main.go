package main

import (
	"MonkeyInterpreterByGO/repl"
	"fmt"
	"os"
	user2 "os/user"
)

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Monkey programming language!\n", user.Username)
	fmt.Printf("Free to type in command\n")
	repl.Start(os.Stdin, os.Stdout)
}
