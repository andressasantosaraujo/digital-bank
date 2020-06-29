package costumer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Costumer struct {
	Name string `json:"name"`
	Cpf string `json:"cpf"`
	Birthday string `json:"birthday"`
}

func CreateCostumer(name, cpf, birthday string) Costumer {
	hasCostumer, costumer := getCostumer(cpf)
	if hasCostumer {
		fmt.Println("tenho costumer")
		return costumer
	} else {
		fmt.Println("não tenho costumer")
		newCostumer := new(Costumer)
		newCostumer.Name = name
		newCostumer.Cpf = cpf
		newCostumer.Birthday = birthday

		setCostumer(*newCostumer)
		return *newCostumer
	}
}

func setCostumer(costumer Costumer) {
	costumersFile, _ := os.OpenFile(`./files/costumers.txt`, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	costumerJson ,_ := json.Marshal(costumer)
	fmt.Println("criei o costumer")
	costumersFile.WriteString(string(costumerJson) + "\n")
}

func getCostumer(cpf string) (bool, Costumer){
	var costumers []Costumer
	costumersFile,_ := os.Open(`./files/costumers.txt`)
	costumersReader := bufio.NewReader(costumersFile)
	for{
		costumer, err := costumersReader.ReadString('\n')
		var costumerJson Costumer
		erro := json.Unmarshal([]byte(strings.TrimSpace(costumer)), &costumerJson)
		if erro != nil {
			fmt.Println(erro)
		}
		costumers = append(costumers, costumerJson)
		if err == io.EOF {
			break
		}
	}
	for _, costumer := range costumers {
		fmt.Println(costumer)
		if costumer.Cpf == cpf{
			return true, costumer
		}
	}
	return false, Costumer{}
}
