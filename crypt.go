package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

const defaultPasswordLength = 20

type crypt struct {
	Name       string
	Categories []category
	password   []byte // password for encryption is not stored in encrypted file. Not exported
}

type category struct {
	Name     string
	Accounts []account
}

type account struct {
	Name     string
	Website  string
	Username string
	Password []byte
	Notes    string
}

func (c *crypt) move(command []string) {
	if len(command) == 3 {
		catIndex, accIndex, accountFound := c.findAccount(command[1])
		if accountFound {
			accountCopy := c.Categories[catIndex].Accounts[accIndex]
			c.deleteAccount(command[1])
			c.addAccount(accountCopy, command[2])
			fmt.Println(command[1], " moved to ", command[2])
		} else {
			fmt.Println("Unable to move ", command[1], " Account not found.")
		}

	} else {
		fmt.Println("Inccorrect move arguments. Please type 'help' for more information.")
	}
}

func newCrypt(cryptName string) {
	c := crypt{Name: cryptName, Categories: []category{}}
	c.password = getPass()
	processCryptCommands(&c)
}

func (c *crypt) closeCrypt() {

	var buffer bytes.Buffer            // Stand-in for a network connection
	encoder := gob.NewEncoder(&buffer) // create encoder
	err := encoder.Encode(c)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	encodedCrypt := buffer.Bytes()

	encryptedData, err := encrypt(c.password, encodedCrypt)
	err = ioutil.WriteFile(path.Join("./crypts/", c.Name), encryptedData, 0644)
	check(err)

	fmt.Println("Crypt Saved.")
	c.wipeCrypt()
	os.Exit(0)
}

func openCrypt(fileName string) {
	password := getPass()
	encryptedFileData, err := ioutil.ReadFile(path.Join("./crypts/", fileName))
	if err != nil {
		fmt.Println("Unable to open", fileName)
		return
	}
	decryptedData, err := decrypt(password, encryptedFileData)
	if err != nil {
		fmt.Println("Failed to open crypt. Wrong password.")
		return
	}

	reader := bytes.NewReader(decryptedData)
	decoder := gob.NewDecoder(reader)
	// Decode (receive) the value.
	var c crypt
	err = decoder.Decode(&c)
	c.password = password
	if err != nil {
		log.Fatal("decode error:", err)
	}

	processCryptCommands(&c)
}

// FindAccount returns the category index location and the account index location
func (c *crypt) findAccount(accountName string) (int, int, bool) {
	for catIndex, category := range c.Categories {
		for accIndex, account := range category.Accounts {
			if account.Name == accountName {
				return catIndex, accIndex, true
			}
		}
	}
	return -1, -1, false
}

func (c *crypt) deleteAccount(accountName string) {

	catIndex, accIndex, accountFound := c.findAccount(accountName)

	if accountFound {
		c.Categories[catIndex].Accounts = append(c.Categories[catIndex].Accounts[:accIndex], c.Categories[catIndex].Accounts[accIndex+1:]...)
		if len(c.Categories[catIndex].Accounts) < 1 {
			c.deleteCategory(catIndex)
		}
	} else {
		fmt.Println("Unable to delete ", accountName, ". Account not found.")
	}

}

func (c *crypt) deleteCategory(catIndex int) {
	c.Categories = append(c.Categories[:catIndex], c.Categories[catIndex+1:]...)
}

// createAccount creates an account instance. The account fields are populated by user input.
func (c *crypt) newAccount(accountName, category string) {
	//TODO scanln only scans line before space
	var account account

	account.Name = accountName
	account.Website = getUserInput("Website: ")
	account.Username = getUserInput("Username: ")
	account.Notes = getUserInput("Notes: ")
	makeNewPassword := askUser("Generate new password? [Y/N]: ")
	if makeNewPassword {
		account.Password = createPassword(defaultPasswordLength)
	} else {
		account.Password = getPass()
	}
	c.addAccount(account, category)
}

func (c *crypt) addAccount(newAccount account, categoryName string) {

	var categoryExists bool
	for i := range c.Categories {
		if c.Categories[i].Name == categoryName {
			categoryExists = true
			c.Categories[i].Accounts = append(c.Categories[i].Accounts, newAccount)
		}
	}
	if !categoryExists {
		category := category{Name: categoryName, Accounts: []account{newAccount}}
		c.Categories = append(c.Categories, category)
	}
}

// func (c *crypt) copy(command []string) {
// 	switch command[1] {
// 	case "-p":
// 		c.copyPassword(command[2])
// 	case "-u":
// 		c.copyUsername(command[2])
// 	}
// }

// func (c *crypt) copyPassword(accountName string) {
// 	catIndex, accIndex, accountFound := c.findAccount(accountName)
// 	if accountFound {
// 		writeToClipboard(string(c.Categories[catIndex].Accounts[accIndex].Password))
// 	} else {
// 		fmt.Println(accountName, " not found.")
// 	}
// }

// func (c *crypt) copyUsername(accountName string) {
// 	catIndex, accIndex, accountFound := c.findAccount(accountName)
// 	if accountFound {
// 		writeToClipboard(c.Categories[catIndex].Accounts[accIndex].Username)
// 	} else {
// 		fmt.Println(accountName, " not found.")
// 	}
// }

func (c *crypt) changePassword(accountName string, passwordLength int) {
	catIndex, accIndex, accountFound := c.findAccount(accountName)
	if accountFound {
		if passwordLength == 0 {
			passwordLength = len(c.Categories[catIndex].Accounts[accIndex].Password)
		}
		c.Categories[catIndex].Accounts[accIndex].Password = createPassword(passwordLength)
	} else {
		fmt.Println(accountName, " not found.")
	}
}

func (c *crypt) wipeCrypt() {
	wipeByteSlice(c.password)
	for i := range c.Categories {
		for j := range c.Categories[i].Accounts {
			wipeByteSlice(c.Categories[i].Accounts[j].Password)
		}
	}

}

func wipeByteSlice(slice []byte) {
	for i := range slice {
		slice[i] = 0
	}
}
