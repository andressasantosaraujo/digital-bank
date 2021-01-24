package main

import (
	"bufio"
	"fmt"
	"os"

	"./account"
	"./customer"
)

func main() {
	introduce()
}

func introduce() {
	fmt.Println(`Welcome to the digital bank!`)
	fmt.Println(`Enter the desired option:`)
	fmt.Println(`0 - Access my account.`)
	fmt.Println(`1 - Create a new account.`)
	fmt.Println(`2 - Exit`)
	var option int
	fmt.Scan(&option)

	switch option {
	case 0:
		login()
	case 1:
		signin()
	case 2:
		os.Exit(0)
	}
}

func getStrIO(name string) string {
	var value string

	fmt.Printf("Enter your %s: ", name)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		value = scanner.Text()

	}
	return value
}

func signin() {
	var accountGot account.Account
	fmt.Println(`***************Create customer***************`)
	cpf := getStrIO("CPF")

	_, err := customer.GetCustomer(cpf)
	if err == nil {
		fmt.Println(`Existing user, log in using the login option.`)
		introduce()
	}

	name := getStrIO("name")
	birthday := getStrIO("birthday")

	err = customer.CreateCustomer(name, cpf, birthday)
	if err != nil {
		fmt.Println(err)
		introduce()
	}

	err, accountGot = account.CreateAccount(cpf)
	if err != nil {
		fmt.Println(err)
		introduce()
	}

	chooseOption(accountGot.Cpf)
}

func login() {
	fmt.Println(`***************Log in***************`)
	cpf := getStrIO("CPF")

	err, account := account.GetAccount(cpf)
	if err != nil {
		fmt.Println(`Non-existent user, sign in.`)
		introduce()
	}

	chooseOption(account.Cpf)
}

func handle_error(err error) {
	if err != nil {
		introduce()
	}
}

func chooseOption(cpf string) {
	var operation string
	fmt.Println(`Welcome to the digital bank!`)
	fmt.Println(`Choose the banking operation:`)
	fmt.Println(`0 - Print balance`)
	fmt.Println(`1 - Cash deposit`)
	fmt.Println(`2 - Withdraw money`)
	fmt.Println(`3 - Transfer money`)
	fmt.Println(`4 - Go back`)
	fmt.Println(`5 - Exit`)
	fmt.Scan(&operation)

	switch operation {
	case `0`:
		err := account.PrintBalance(cpf)
		handle_error(err)
	case `1`:
		err := account.Deposit(cpf)
		handle_error(err)
	case `2`:
		err := account.Withdraw(cpf)
		handle_error(err)
	case `3`:
		err := account.Transfer(cpf)
		handle_error(err)
	case `4`:
		introduce()
	case `5`:
		os.Exit(0)
	}

	chooseOption(cpf)
}
