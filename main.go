package main

import (
	"fmt"
	"ridersenis/bank/fileOps"
)

const balanceFileName = "balance.txt"

func print(text string) {
	fmt.Println(text)
}

func showOptions() {
	print("What would you like to do ?")
	print("1. Check balance")
	print("2. Deposit money")
	print("3. Witdrawl money")
	print("4. Exit")
}
func choicesControl(balance float64) {
	var choice int
	fmt.Println("Welcome to Go Bank Application")
	for {
		showOptions()
		fmt.Print("Your choice:")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your current balance is: ", balance)
		case 2:
			var depositAmount float64
			fmt.Print("Deposit amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				print("Deposit must be greater than 0 !")
				continue
			}
			balance = balance + depositAmount
			fileOps.WriteFloatToFile(balance, balanceFileName)
			fmt.Println("Your Current Balance: ", balance)
		case 3:
			var witdrawlAmount float64
			fmt.Print("Witdrawl amount: ")
			fmt.Scan(&witdrawlAmount)
			if witdrawlAmount > balance {
				print("No enough amount to perform this operation !")
				continue
			} else if witdrawlAmount <= 0 {
				print("Mount must be greater than 0 ")
				continue
			}
			balance = balance - witdrawlAmount
			fileOps.WriteFloatToFile(balance, balanceFileName)
			print("Witdrawl was sucessfully ")
			fmt.Print("Your Current Balance: ", balance)
		default:
			print("Sesion Terminada!")
			return

		}
	}

}
func main() {
	balance, err := fileOps.ReadFloatFromFile(balanceFileName)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------------------------------------------")
	}
	choicesControl(balance)
}
