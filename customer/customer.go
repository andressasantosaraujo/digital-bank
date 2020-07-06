package customer

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strings"
)

type Customer struct {
	Name string `json:"name"`
	Cpf string `json:"cpf"`
	Birthday string `json:"birthday"`
}

func CreateCustomer(name, cpf, birthday string) error {
	_, err  := getCustomer(cpf)
	if err != nil {
		newCostumer := Customer{name,cpf,birthday}
		err = setCustomer(newCostumer)
	}
	return err
}

func setCustomer(costumer Customer) error {
	costumersFile, err := os.OpenFile(`./files/costumers.txt`, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err == nil {
		costumerJson , err := json.Marshal(costumer)
		if err == nil {
			costumersFile.WriteString(string(costumerJson) + "\n")
			return nil
		}
	}
	return err
}

func getCustomer(cpf string) (Customer, error){
	costumersFile, err := os.Open(`./files/costumers.txt`)
	if costumersFile != nil  {
		costumersReader := bufio.NewReader(costumersFile)
		for err != io.EOF {
			var costumerJson Customer
			costumer, err := costumersReader.ReadString('\n')
			err = json.Unmarshal([]byte(strings.TrimSpace(costumer)), &costumerJson)
			if err == nil {
				if costumerJson.Cpf == cpf{
					return costumerJson, nil
				}
			}
		}
	}
	return Customer{}, err
}
