package main

import "fmt"

func print(text string) {
	fmt.Println(text)
}
func getUserBalance() float64 {
	const balance float64 = 1000.0
	return balance
}
func choicesControl(balance float64) {
	var choice int
	print("Welcome to Go Bank Application")
	for {
		print("What would you like to do ?")
		print("1. Check balance")
		print("2. Deposit money")
		print("3. Witdrawl money")
		print("4. Exit")
		fmt.Print("Your choice:")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your current balance is: ", balance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("Deposit amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				print("Deposit must be greater than 0 !")
				continue
			}
			balance = balance + depositAmount
			fmt.Println("Your Current Balance: ", balance)
		} else if choice == 3 {
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
			print("Witdrawl was sucessfully ")
			fmt.Print("Your Current Balance: ", balance)
		} else {
			print("Sesion Terminada!")
			break
		}
	}

}
func main() {
	balance := getUserBalance()
	choicesControl(balance)
}
