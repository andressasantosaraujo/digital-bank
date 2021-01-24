package customer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Customer struct {
	Name     string `json:"name"`
	Cpf      string `json:"cpf"`
	Birthday string `json:"birthday"`
}

func CreateCustomer(name, cpf, birthday string) error {
	_, err := GetCustomer(cpf)
	if err != nil {
		newcustomer := Customer{name, cpf, birthday}
		err = setCustomer(newcustomer)
	}
	return err
}

func setCustomer(customer Customer) error {
	customersFile, err := os.OpenFile(`./files/customers.txt`, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	customerJson, err := json.Marshal(customer)
	if err != nil {
		return err
	}
	_, err = customersFile.WriteString(string(customerJson) + "\n")
	if err != nil {
		return err
	}
	return nil
}

func GetCustomer(cpf string) (Customer, error) {
	var customerJson Customer
	customerFile, err := os.Open(`./files/customers.txt`)
	if err != nil {
		return Customer{}, err
	}
	customerReader := bufio.NewReader(customerFile)
	for {
		customer, err := customerReader.ReadString('\n')
		_ = json.Unmarshal([]byte(strings.TrimSpace(customer)), &customerJson)
		if customerJson.Cpf == cpf {
			customerFile.Close()
			return customerJson, nil
		}
		if err == io.EOF {
			break
		}
	}
	customerFile.Close()
	err = fmt.Errorf(`Error to get Customer`)
	return Customer{}, err
}
