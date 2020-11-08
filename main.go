package main

import (
	"fmt"

	"./account"
	"./customer"
)

func main() {
	introduce()
}

func introduce() {
	var option int
	fmt.Println(`Welcome to the digital bank!`)
	fmt.Println(`Enter the desired option:`)
	fmt.Println(`1 - Access my account.`)
	fmt.Println(`2 - Create a new account.`)
	fmt.Scan(&option)

	switch option {
	case 1:
		login()
	case 2:
		signin()
	}
}

func login() {
	var accountNumber string
	fmt.Println(`Enter your account number:`)
	fmt.Scan(&accountNumber)
	// account.GetAccount(accountNumber)
}

func signin() {
	var name string
	var cpf string
	var birthday string
	var numberAccount string
	var operation int

	fmt.Println(`***************Create current customer***************`)
	fmt.Println(`Type your name:`)
	fmt.Scan(&name)
	fmt.Println(`Type your CPF:`)
	fmt.Scan(&cpf)
	fmt.Println(`Type your birthday:`)
	fmt.Scan(&birthday)

	err := customer.CreateCustomer(name, cpf, birthday)
	if err == nil {
		fmt.Println(`Type your number account:`)
		fmt.Scan(&numberAccount)
		fmt.Println(`Type your operation:`)
		fmt.Scan(&operation)
		err = account.CreateAccount(operation, cpf, numberAccount)
	}

}

func chooseOption(accountNumber string) {
	var operation string
	fmt.Println(`Welcome to the digital bank!`)
	fmt.Println(`Choose the banking operation:`)
	fmt.Println(`1 - Open an account`)
	fmt.Println(`2 - Print balance`)
	fmt.Println(`3 - Cash deposit`)
	fmt.Println(`4 - Withdraw money`)
	fmt.Println(`5 - Transfer money`)
	fmt.Scan(&operation)

	switch operation {
	case `1`:
		account.CreateAccount(1, accountNumber, "")
	case `2`:
		account.Deposit(accountNumber)
	case `3`:
		account.Withdraw(accountNumber)
	case `4`:
		account.PrintBalance(accountNumber)
	case `5`:
		account.Transfer(accountNumber)
	}
}
