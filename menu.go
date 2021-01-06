package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func getCommand(prompt string) []string {
	commandScanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	commandScanner.Scan()
	command := commandScanner.Text()
	return strings.Split(command, " ")
}

func printStartingCommands() {
	fmt.Println("\n\tcreate \033[4mCRYPT_NAME\033[0m",
		"\n\t\t create a new crypt",
		"\n\n\topen \033[4mCRYPT_NAME\033[0m",
		"\n\t\topen a crypt")
}

func printCommands() {

	fmt.Println("\n\tshow [ACCOUNT]\n",
		"\n\tnew \033[4mACCOUNT\033[0m \033[4mCATEGORY\033[0m\n",
		"\n\tdelete \033[4mACCOUNT\033[0m\n",
		"\n\tmove \033[4mACCOUNT\033[0m \033[4mDESTINATION\033[0m\n",
		"\n\tclose")
}

func getPass() []byte {
	fmt.Print("Password: ")
	password, err := terminal.ReadPassword(0)
	fmt.Println()
	if err != nil {
		fmt.Print("Unable to silence input. Please re-enter password,", "\nPassword: ")
		fmt.Scanln(&password) // TODO check if valid
	}
	return password
}

func getUserInput(message string) string {
	fmt.Print(message)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func askUser(question string) bool {
	for {
		response := getUserInput(question)
		switch strings.ToLower(response) {
		case "y":
			return true
		case "n":
			return false
		default:
			fmt.Println("Please enter 'Y' for yes or 'N' for no.")
		}
	}
}

func processCommands() {
	for {
		command := getCommand("Crypt: ")
		switch command[0] {
		case "exit":
			os.Exit(0)
		case "help":
			printStartingCommands()
		case "open":
			if len(command) > 1 {
				openCrypt(command[1])
			} else {
				fmt.Println("Crypt name not specified.")
			}

		case "create":
			if len(command) > 1 {
				newCrypt(command[1])
			} else {
				fmt.Println("Crypt name not specified.")
			}

		}
	}
}

func processCryptCommands(c *crypt) {
	for {
		command := getCommand(c.Name + ": ")
		switch command[0] {
		case "show":
			showCrypt(c, command)
		case "new":
			if len(command) == 3 {
				c.newAccount(command[1], command[2])
			} else {
				fmt.Println("Missing command arguments. Usage: new ACCOUNT CATEGORY")
			}
		case "delete":
			if len(command) > 1 {
				c.deleteAccount(command[1])
				fmt.Println(command[1], " deleted.")
			} else {
				fmt.Println("No account specified for deletion.")
			}
		case "cp":
			passwordLength := 0
			if len(command) < 2 {
				fmt.Println("Command arguments not valid.")
			}
			if len(command) > 2 {
				var err error
				passwordLength, err = strconv.Atoi(command[2])
				if err != nil {
					fmt.Println("Command arguments not valid.")
				}
			}
			c.changePassword(command[1], passwordLength)
		case "copy":
			//c.copy(command)
		case "move":
			c.move(command)
		case "close":
			c.closeCrypt()
			os.Exit(0)
		case "help":
			printCommands()
		default:
			fmt.Println("Command not recognized. Please enter 'help' for a list of commands.")
		}
	}
}
