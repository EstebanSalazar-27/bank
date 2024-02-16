package main

import "fmt"

func print(text string) {
	fmt.Println(text)
}
func getUserBalance() float64 {
	const balance float64 = 1000.0
	return balance
}
func choicesControl(choice int, balance float64) {
	if choice == 1 {
		fmt.Print("Your current balance is: ", balance)
	} else if choice == 2 {
		var depositAmount float64
		fmt.Print("Deposit amount: ")
		fmt.Scan(&depositAmount)
		balance = balance + depositAmount
		fmt.Print("Your Current Balance: ", balance)
	} else if choice == 3 {
		var witdrawlAmount float64
		fmt.Print("Witdrawl amount: ")
		fmt.Scan(&witdrawlAmount)
		balance = balance - witdrawlAmount
		print("Witdrawl was sucessfully ")
		fmt.Print("Your Current Balance: ", balance)
	} else {
		print("Sesion Terminada!")
	}
}
func main() {
	balance := getUserBalance()
	print("Welcome to Go Bank Application")
	print("What would you like to do ?")
	print("1. Check balance")
	print("2. Deposit money")
	print("3. Witdrawl money")
	print("4. Exit")

	var choice int
	fmt.Print("Your choice:")
	fmt.Scan(&choice)
	choicesControl(choice, balance)

}
