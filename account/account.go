package account

import (
	"../customer"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Account struct {
	Number    string            `json:"number"`
	Money     float64           `json:"money"`
	Operation int               `json:"operation"`
	Costumer  customer.Customer `json:"customer"`
}

func GetAccount(accountNumber string) (Account, error) {
	accountsFile, err := os.Open(`./files/account.txt`)
	if accountsFile != nil  {
		accountReader := bufio.NewReader(accountsFile)
		for err != io.EOF {
			var accountJson Account
			account, err := accountReader.ReadString('\n')
			err = json.Unmarshal([]byte(strings.TrimSpace(account)), &accountJson)
			if err == nil {
				if accountJson.Number == accountNumber{
					return accountJson, nil
				}
			}
		}
	}
	return Account{}, err
}

func SetAccount(account Account)  error {
	accountsFile, err := os.OpenFile(`./files/accounts.txt`, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err == nil {
		accountJson , err := json.Marshal(account)
		if err == nil {
			accountsFile.WriteString(string(accountJson) + "\n")
			return nil
		}
	}
	return err
}

func CreateAccount(customer customer.Customer, number string, operation int) (Account, error) {
	account, err  := GetAccount(number)
	if err != nil {
		newAccount := Account{number,0,operation, customer}
		err = SetAccount(newAccount)
		return newAccount, err
	}
	return account, err
}

func Deposit(accountNumber string) {
	var value float64
	fmt .Println(`***************Cash deposit***************`)
	fmt.Println(`Enter the value to be deposited:`)
	fmt.Scan(&value)
}

func Withdraw(accountNumber string)  {
	fmt.Println(`***************Withdraw money***************`)

}

func PrintBalance(accountNumber string)  {
	fmt.Println(`***************Withdraw money***************`)

}

func Transfer(accountNumber string) {
	fmt.Println(`***************Transfer money between accounts1***************`)

}