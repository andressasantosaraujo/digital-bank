package main

import (
	"./account"
	"./costumer"
	"fmt"
)

func main()  {
	introduce()
}

func introduce(){
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
	account.GetAccount()
}

func signin() {
	var name string
	var cpf string
	var birthday string

	fmt.Println(`***************Create current costumer***************`)
	fmt.Println(`Type your name:`)
	fmt.Scan(&name)
	fmt.Println(`Type your CPF:`)
	fmt.Scan(&cpf)
	fmt.Println(`Type your birthday:`)
	fmt.Scan(&birthday)

	costumer := costumer.CreateCostumer(name, cpf, birthday)
	fmt.Println(costumer)
	accountObj := account.CreateAccount()
	fmt.Println(accountObj.Number)
	chooseOption(accountObj.Number)
}

func chooseOption(accountNumber string) {
	var operation string
	fmt.Println(`Choose the banking operation:`)
	fmt.Println(`1 - Open an account`)
	fmt.Println(`2 - Print balance`)
	fmt.Println(`3 - Cash deposit`)
	fmt.Println(`4 - Withdraw money`)
	fmt.Println(`5 - Transfer money`)
	fmt.Scan(&operation)

	switch operation {
	case `1`:
		account.Deposit(accountNumber)
	case `2`:
		account.Withdraw(accountNumber)
	case `3`:
		account.Transfer(accountNumber)
	case `4`:
		account.PrintBalance(accountNumber)
	}
}