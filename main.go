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
	}
	if choice == 2 {
		var deposit float64
		fmt.Print("Deposit: ")
		fmt.Scan(&deposit)
		balance = balance + deposit
		fmt.Print("Your Current Balance: ", balance)
	}
	if choice == 3 {
		var withdraw float64
		fmt.Print("amount: ")
		fmt.Scan(&withdraw)
		balance = balance - withdraw
		print("Witdraw was sucessfully ")
		fmt.Print("Your Current Balance: ", balance)
	}
}
func main() {
	balance := getUserBalance()
	print("Welcome to Go Bank Application")
	print("What would you like to do ?")
	print("1. Check balance")
	print("2. Deposit money")
	print("3. Withdraw money")
	print("4. Exit")

	var choice int
	fmt.Print("Your choice:")
	fmt.Scan(&choice)
	choicesControl(choice, balance)

}
