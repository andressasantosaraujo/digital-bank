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

func getAccount(number string) (Account, error) {
	var accountJson Account
	accountFile, err := os.Open(`./files/accounts.txt`)
	if err != nil {
		return Account{}, err
	}
	accountReader := bufio.NewReader(accountFile)
	for {
		account, err := accountReader.ReadString('\n')
		_ = json.Unmarshal([]byte(strings.TrimSpace(account)), &accountJson)

		if accountJson.Number == number {
			accountFile.Close()
			return accountJson, nil
		}
		if err == io.EOF {
			break
		}
	}
	accountFile.Close()
	err = fmt.Errorf(`Error to get Account`)
	return Account{}, err
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

func CreateAccount(operation int, cpf string, number string) error {
	_, err := getAccount(number)
	if err != nil {
		newAccount := Account{number, 0, operation, cpf}
		err = setAccount(newAccount)
	}
	return err
}

func Deposit(accountNumber string) {
	var value float64
	fmt.Println(`***************Cash deposit***************`)
	fmt.Println(`Enter the value to be deposited:`)
	fmt.Scan(&value)
}

func Withdraw(accountNumber string) {
	fmt.Println(`***************Withdraw money***************`)

}

func PrintBalance(accountNumber string) {
	fmt.Println(`***************Withdraw money***************`)

}

func Transfer(accountNumber string) {
	fmt.Println(`***************Transfer money between accounts1***************`)

}
