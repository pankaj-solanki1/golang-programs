package main

import (
	"fmt"
	"go-programs/RLHF/golang_random/28-11-24/389103/turn2/modelB/bankaccount"
)

func main() {
	// Create a new bank account for Alice with an initial balance of $100.
	account, err := bankaccount.NewBankAccount("Alice", 100.0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Account Holder: %s\n", account.Name())
	fmt.Printf("Initial Balance: $%.2f\n", account.Balance())

	// Deposit $50 into the account.
	err = account.Deposit(50.0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Balance after Deposit: $%.2f\n", account.Balance())

	// Withdraw $30 from the account.
	err = account.Withdraw(30.0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Balance after Withdrawal: $%.2f\n", account.Balance())

	// Attempt to withdraw an invalid amount.
	err = account.Withdraw(200.0)
	fmt.Println(err)

	// Print transaction history.
	for _, transaction := range account.TransactionHistory() {
		fmt.Printf("%s: $%.2f - %s\n", transaction.Type, transaction.Amount, transaction.Description)
	}
}
