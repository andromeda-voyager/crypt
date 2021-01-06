package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		os.Exit(0)
	}
}

func main() {

	fmt.Println("\t-  PasswordCrypt Password Manager  -", "\n\nEnter 'help' for a list of commands.")
	processCommands()

}
