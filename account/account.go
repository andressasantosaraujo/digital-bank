package account

import (
	"../costumer"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Account struct {
	Number    string `json:"number"`
	Money     float64 `json:"money"`
	Operation int `json:"operation"`
	Costumer costumer.Costumer `json:"costumer"`
}

func GetAccount()  {
	var accounts []string
	accountsFile,_ := os.Open(`./files/accounts.txt`)
	accountReader := bufio.NewReader(accountsFile)
	for {
		account, err := accountReader.ReadString('\n')
		account = strings.TrimSpace(account)
		accounts = append(accounts, account)
		if err == io.EOF {
			break
		}
	}
	for i, v := range accounts{
		fmt.Println(i, v)
	}
}

func SetAccount(account Account)  {
	accountFile, _ := os.OpenFile(`./files/accounts.txt`, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	obj,_ := json.Marshal(account)
	accountFile.WriteString(string(obj))
}

func CreateAccount() Account {
	var accountType int
	fmt.Println(`***************Create current account***************`)
	fmt.Println(`Choose your account type:`)
	fmt.Println(`1 - Savings account.`)
	fmt.Println(`2 - Current account.`)
	fmt.Scan(&accountType)

	newAccount := new(Account)
	newAccount.Operation = accountType
	newAccount.Money = 0
	newAccount.Number = `0`

	SetAccount(*newAccount)
	return *newAccount
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