package main

import "fmt"

const (
	branch     = "├──────"
	twig       = "│\t├──────"
	leaf       = "│\t└──────"
	lastBranch = "└──────"
	lastTwig   = "\t├──────"
	lastLeaf   = "\t└──────"
)

func (c *crypt) showAccount(accountName string, mask bool) {

	catIndex, accIndex, accountFound := c.findAccount(accountName)
	if accountFound {
		fmt.Println("\n\tAccount:", c.Categories[catIndex].Accounts[accIndex].Name)
		fmt.Println("\tWebsite:", c.Categories[catIndex].Accounts[accIndex].Website)
		fmt.Println("\tUsername:", c.Categories[catIndex].Accounts[accIndex].Username)

		if mask {
			fmt.Println("\tPassword: *******")
			fmt.Println("\tNotes: *******", c.Categories[catIndex].Accounts[accIndex].Notes)
		} else {
			fmt.Println("\tPassword:", string(c.Categories[catIndex].Accounts[accIndex].Password))
			fmt.Print("\tNotes:", c.Categories[catIndex].Accounts[accIndex].Notes, "\n\n")
		}

	}
}

func showCrypt(crypt *crypt, command []string) {

	if len(command) > 1 {
		crypt.showAccount(command[1], false)
	} else {
		fmt.Println(crypt.Name)
		if len(crypt.Categories) > 0 {
			for _, category := range crypt.Categories[:len(crypt.Categories)-1] {
				fmt.Println(branch, category.Name)
				printAccounts(category.Accounts)
			}
			printLastCategory(crypt.Categories[len(crypt.Categories)-1])
		}
	}

}

func printAccounts(accounts []account) {
	for _, account := range accounts[:len(accounts)-1] {
		fmt.Println(twig, account.Name)
	}
	fmt.Println(leaf, accounts[len(accounts)-1].Name)
}

func printLastCategory(category category) {
	fmt.Println(lastBranch, category.Name)
	for _, account := range category.Accounts[:len(category.Accounts)-1] {
		fmt.Println(lastTwig, account.Name)
	}
	fmt.Println(lastLeaf, category.Accounts[len(category.Accounts)-1].Name)
}
