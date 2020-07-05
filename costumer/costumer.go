package costumer

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strings"
)

type Costumer struct {
	Name string `json:"name"`
	Cpf string `json:"cpf"`
	Birthday string `json:"birthday"`
}

func CreateCostumer(name, cpf, birthday string) error {
	_, err  := getCostumer(cpf)
	if err != nil {
		newCostumer := Costumer{name,cpf,birthday}
		err = setCostumer(newCostumer)
	}
	return err
}

func setCostumer(costumer Costumer) error {
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

func getCostumer(cpf string) (Costumer, error){
	costumersFile, err := os.Open(`./files/costumers.txt`)
	if costumersFile != nil  {
		costumersReader := bufio.NewReader(costumersFile)
		for err != io.EOF {
			var costumerJson Costumer
			costumer, err := costumersReader.ReadString('\n')
			err = json.Unmarshal([]byte(strings.TrimSpace(costumer)), &costumerJson)
			if err == nil {
				if costumerJson.Cpf == cpf{
					return costumerJson, nil
				}
			}
		}
	}
	return Costumer{}, err
}
