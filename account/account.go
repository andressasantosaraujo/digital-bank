package account

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Account struct {
	Number    string  `json:"number"`
	Money     float64 `json:"money"`
	Operation int     `json:"operation"`
	Cpf       string  `json:"cpf"`
}

func setAccount(account Account) error {
	accountFile, err := os.OpenFile(`./files/accounts.txt`, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	accountJson, err := json.Marshal(account)
	if err != nil {
		return err
	}
	accountFile.WriteString(string(accountJson) + "\n")
	return nil
}

func CreateAccount(operation int, cpf string, number string) (error, Account) {
	var newAccount Account
	err, _ := GetAccount(cpf)
	if err != nil {
		newAccount = Account{number, 0, operation, cpf}
		err = setAccount(newAccount)
	}
	return err, newAccount
}

func GetAccount(cpf string) (error, Account) {
	var accountJson Account
	accountFile, err := os.Open(`./files/accounts.txt`)
	if err != nil {
		return err, Account{}
	}
	accountReader := bufio.NewReader(accountFile)
	for {
		account, err := accountReader.ReadString('\n')
		_ = json.Unmarshal([]byte(strings.TrimSpace(account)), &accountJson)

		if accountJson.Cpf == cpf {
			accountFile.Close()
			return nil, accountJson
		}
		if err == io.EOF {
			break
		}
	}
	accountFile.Close()
	err = fmt.Errorf(`Error to get Account`)
	return err, Account{}
}

func checkMoney(money float64) (error, float64) {
	var value float64
	fmt.Println(`Enter the transaction amount:`)
	fmt.Scan(&value)
	if err := fmt.Errorf(`err`); money <= value {
		return err, 0
	}
	return nil, value
}

func getMoney(account Account, value float64) {
	account.Money -= value
}
func setMoney(account Account, value float64) {
	account.Money += value
}

func Deposit(account Account) {
	fmt.Println(`***************Cash deposit***************`)
	var value float64
	fmt.Println(`Enter the transaction amount:`)
	fmt.Scan(&value)
	setMoney(account, value)
}

func Withdraw(account Account) {
	var value float64
	fmt.Println(`***************Cash deposit***************`)
	err, value := checkMoney(account.Money)
	if err != nil {
		fmt.Println(`Insufficient balance!`)
	}
	getMoney(account, value)
}

func PrintBalance(account Account) {
	fmt.Println(`***************Print balance***************`)
	fmt.Println(account.Money)
}

func Transfer(account Account) {
	fmt.Println(`***************Transfer money between accounts***************`)
	var value float64
	var destinyCPF string
	err, value := checkMoney(account.Money)
	if err != nil {
		fmt.Println(`Insufficient balance!`)
	}
	fmt.Println(`Enter the destination account for the transfer:`)
	fmt.Scan(&destinyCPF)
	err, destinyAccount := GetAccount(destinyCPF)
	getMoney(account, value)
	setMoney(destinyAccount, value)
}
