package main

import (
	"fmt"
)

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.Balance += amount
		fmt.Printf("Deposited: %.2f\n", amount)
	} else {
		fmt.Println("Invalid deposit amount.")
	}
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount > 0 && amount <= b.Balance {
		b.Balance -= amount
		fmt.Printf("Withdrawn: %.2f\n", amount)
	} else if amount > b.Balance {
		fmt.Println("Insufficient funds.")
	} else {
		fmt.Println("Invalid withdrawal amount.")
	}
}

func (b BankAccount) GetBalance() {
	fmt.Printf("Current balance: %.2f\n", b.Balance)
}

func Transaction(account *BankAccount, transactions []float64) {
	for _, transaction := range transactions {
		if transaction > 0 {
			account.Deposit(transaction)
		} else {
			account.Withdraw(-transaction)
		}
	}
}

func main() {
	var account BankAccount
	account.AccountNumber = "12345678"
	account.HolderName = "Nursultan Serikuly"
	account.Balance = 100

	var action string
	var amount float64

	for {
		fmt.Println("Choose an action:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Get balance")
		fmt.Println("4. Exit")
		fmt.Print("Enter choice: ")
		fmt.Scanln(&action)

		switch action {
		case "1":
			fmt.Print("Enter deposit amount: ")
			fmt.Scanln(&amount)
			account.Deposit(amount)

		case "2":
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scanln(&amount)
			account.Withdraw(amount)

		case "3":
			account.GetBalance()

		case "4":
			fmt.Println("Exiting the system.")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
