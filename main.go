package main

import (
	"fmt"
)

func main()  {
	chooseOption()
}

func chooseOption() {
	var operation string
	fmt.Println(`Welcome to the digital bank!`)
	fmt.Println(`Choose the banking operation:`)
	fmt.Println(`1 - Open an account`)
	fmt.Println(`2 - Print balance`)
	fmt.Println(`3 - Cash deposit`)
	fmt.Println(`4 - Withdraw money`)
	fmt.Println(`5 - Transfer money`)
	fmt.Scan(&operation)

	switch operation {
	case `1`:
		accountCreate()
	case `2`:
		deposit()
	case `3`:
		withdraw()
	case `4`:
		printBalance()
	case `5`:
		transfer()
	}
}

func accountCreate()  {

}

func deposit() {
	
}

func withdraw()  {
	
}

func printBalance()  {
	
}

func transfer() {

}