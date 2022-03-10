package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		help()
	} else {
		// handle each command
	}
}

func help() {
	fmt.Println("dfm init -- create an empty repo to fill with dotfiles")
	fmt.Println("dfm sync -- update dotfiles with remote changes, and push yours")
	fmt.Println("dfm fetch -- pull a repo and deploy it")
	fmt.Println("dfm deploy -- deploy your current changes to your system")
}
