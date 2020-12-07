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
	fmt.Println(`0 - Access my account.`)
	fmt.Println(`1 - Create a new account.`)
	fmt.Scan(&option)

	switch option {
	case 0:
		login()
	case 1:
		signin()
	}
}

func signin() {
	var name string
	var cpf string
	var birthday string
	var numberAccount string
	var operation int
	var accountGot account.Account

	fmt.Println(`***************Create customer***************`)
	fmt.Println(`Type your CPF:`)
	fmt.Scan(&cpf)
	_, err := customer.GetCustomer(cpf)
	if err == nil {
		fmt.Println(`Existing user, log in using the login option.`)
		introduce()
	}

	fmt.Println(`Type your name:`)
	fmt.Scan(&name)
	fmt.Println(`Type your birthday:`)
	fmt.Scan(&birthday)
	err = customer.CreateCustomer(name, cpf, birthday)

	if err == nil {
		fmt.Println(`Type your number account:`)
		fmt.Scan(&numberAccount)
		err, accountGot = account.GetAccount(numberAccount)
		if err == nil {
			fmt.Println(`Existing account, choose the banking operation:`)
			chooseOption(accountGot)
		}
		fmt.Println(`Type your operation:`)
		fmt.Scan(&operation)
		err, accountGot = account.CreateAccount(operation, cpf, numberAccount)
	}
	chooseOption(accountGot)
}

func login() {
	var cpf string
	fmt.Println(`***************Log in***************`)
	fmt.Println(`Type your CPF:`)
	fmt.Scan(&cpf)
	err, account := account.GetAccount(cpf)
	if err != nil {
		fmt.Println(`Non-existent user, sign in.`)
		signin()
	}
	chooseOption(account)
}

func chooseOption(accountGot account.Account) {
	var operation string
	fmt.Println(`Welcome to the digital bank!`)
	fmt.Println(`Choose the banking operation:`)
	fmt.Println(`0 - Print balance`)
	fmt.Println(`1 - Cash deposit`)
	fmt.Println(`2 - Withdraw money`)
	fmt.Println(`3 - Transfer money`)
	fmt.Println(`4 - Go back`)
	fmt.Scan(&operation)

	switch operation {
	case `0`:
		account.PrintBalance(accountGot)
	case `1`:
		account.Deposit(accountGot)
	case `2`:
		account.Withdraw(accountGot)
	case `3`:
		account.Transfer(accountGot)
	case `4`:
		introduce()
	}

	chooseOption(accountGot)
}
