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
	Number    int     `json:"number"`
	Money     float64 `json:"money"`
	Operation int     `json:"operation"`
	Cpf       string  `json:"cpf"`
}

func setAccount(account Account) error {
	accountFile, err := os.OpenFile(`./files/accounts.txt`, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	accountJSON, err := json.Marshal(account)
	if err != nil {
		return err
	}

	_, err = accountFile.WriteString(string(accountJSON) + "\n")
	if err != nil {
		return err
	}
	return nil
}

func CreateAccount(cpf string) (error, Account) {
	var newAccount Account

	err, accounts := getAccounts()

	if err == nil && len(accounts) >= 1 {
		for _, account := range accounts {
			if account.Cpf == cpf {
				return fmt.Errorf("Account already exists!"), account
			}
		}
	}

	newAccountNumber := len(accounts) + 1
	newAccount = Account{newAccountNumber, 0, getOprAccount(), cpf}

	_ = setAccount(newAccount)

	return nil, newAccount
}

func GetAccount(cpf string) (error, Account) {
	var accountJSON Account

	accountFile, err := os.Open(`./files/accounts.txt`)
	if err != nil {
		return err, Account{}
	}

	accountReader := bufio.NewReader(accountFile)
	for {
		account, err := accountReader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			return err, Account{}
		}

		err = json.Unmarshal([]byte(strings.TrimSpace(account)), &accountJSON)
		if err != nil {
			return err, Account{}
		}

		if accountJSON.Cpf == cpf {
			accountFile.Close()
			return nil, accountJSON
		}
	}

	accountFile.Close()
	err = fmt.Errorf(`Invalid account`)

	return err, Account{}
}

func checkMoney(money float64, value float64) error {
	if err := fmt.Errorf(`Insufficient balance!`); money < value {
		return err
	}

	return nil
}

func getMoney(account Account, value float64) error {
	_, accounts := getAccounts()
	for i := range accounts {
		if account.Cpf == accounts[i].Cpf {
			accounts[i].Money -= value
		}
	}

	err := setAccounts(accounts)
	if err != nil {
		return err
	}

	return nil
}

func setMoney(account Account, value float64) error {
	_, accounts := getAccounts()
	for i := range accounts {
		if account.Cpf == accounts[i].Cpf {
			accounts[i].Money += value
		}
	}

	err := setAccounts(accounts)
	if err != nil {
		return err
	}

	return nil
}

func getOprAccount() int {
	fmt.Println(`Choose the account operation: `)
	fmt.Println(`0 - Current account`)
	fmt.Println(`1 - Save account`)
	fmt.Println(`2 - Business account`)

	var option int
	fmt.Scan(&option)

	return option
}

func getAmountIO() float64 {
	var value float64
	fmt.Println(`Enter the transaction amount:`)
	fmt.Scan(&value)
	return value
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

func Deposit(cpf string) error {
	fmt.Println(`***************Cash deposit***************`)

	err, account := GetAccount(cpf)
	if err != nil {
		fmt.Println(`Invalid CPF!`)
		return err
	}

	value := getAmountIO()

	err = setMoney(account, value)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func Withdraw(cpf string) error {
	fmt.Println(`***************Cash deposit***************`)

	err, account := GetAccount(cpf)
	if err != nil {
		fmt.Println(`Invalid CPF!`)
	}

	value := getAmountIO()
	err = checkMoney(account.Money, value)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = getMoney(account, value)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func Transfer(cpf string) error {
	fmt.Println(`***************Transfer money between accounts***************`)

	err, account := GetAccount(cpf)
	if err != nil {
		fmt.Println(`Invalid CPF!`)
	}

	destinyCPF := getStrIO("destination CPF")

	err, destinyAccount := GetAccount(destinyCPF)
	if err != nil {
		fmt.Println(err)
		return err
	}

	value := getAmountIO()

	err = checkMoney(account.Money, value)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = getMoney(account, value)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = setMoney(destinyAccount, value)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func getAccounts() (error, []Account) {
	var accountJDON Account
	var accounts []Account
	accountFile, err := os.Open(`./files/accounts.txt`)

	if err != nil {
		return err, accounts
	}

	accountReader := bufio.NewReader(accountFile)

	for {
		account, err := accountReader.ReadString('\n')

		if err == io.EOF {
			break
		}

		_ = json.Unmarshal([]byte(strings.TrimSpace(account)), &accountJDON)
		accounts = append(accounts, accountJDON)
	}

	accountFile.Close()
	return err, accounts
}

func setAccounts(accounts []Account) error {
	err := os.Remove("./files/accounts.txt")

	if err != nil {
		return err
	}

	accountFile, err := os.OpenFile(`./files/accounts.txt`, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return err
	}

	for _, account := range accounts {
		accountJSON, err := json.Marshal(account)

		if err != nil {
			return err
		}

		_, err = accountFile.WriteString(string(accountJSON) + "\n")
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func PrintBalance(cpf string) error {
	fmt.Println(`***************Print balance***************`)

	err, account := GetAccount(cpf)
	if err != nil {
		fmt.Println(`Invalid CPF!`)
		return err
	}

	fmt.Println(account.Money)

	return nil
}
