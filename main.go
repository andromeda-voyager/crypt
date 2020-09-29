package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		os.Exit(0)
	}
}

func getCommand() []string {
	commandScanner := bufio.NewScanner(os.Stdin)
	fmt.Print("PasswordCrypt: ")
	commandScanner.Scan()
	command := commandScanner.Text()
	return strings.Split(command, " ")
}

func printStartingCommands() {
	fmt.Println("\n\tcreate CRYPT",
		"\n\t\t create a new crypt",
		"\n\topen CRYPT",
		"\n\t\topen a crypt file")
}

func printCommands() {

	fmt.Println("\n\tshow [ACCOUNT]\n",
		"\n\tnew ACCOUNT CATEGORY\n",
		"\n\tdelete ACCOUNT\n",
		"\n\tmove ACCOUNT DESTINATION\n",
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
		command := getCommand()
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
		command := getCommand()
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

func main() {
	//fmt.Println(createPassword(20))

	fmt.Println("\t-  PasswordCrypt Password Manager  -", "\n\nEnter 'help' for a list of commands.")

	processCommands()

}
