package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	owner   string
	balance float64
}

func NewBankAccount(owner string, initialBalance float64) (*BankAccount, error) {
	if owner == "" {
		return nil, errors.New("владелец счета не может быть пустым")
	}
	if initialBalance < 0 {
		return nil, errors.New("начальный баланс не может быть отрицательным")
	}

	return &BankAccount{
		owner:   owner,
		balance: initialBalance,
	}, nil
}

func (acc *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма пополнения должна быть положительной")
	}
	acc.balance += amount
	return nil
}

func (acc *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма снятия должна быть положительной")
	}
	if amount > acc.balance {
		return errors.New("недостаточно средств на счете")
	}
	acc.balance -= amount
	return nil
}

func (acc *BankAccount) Balance() float64 {
	return acc.balance
}

func (acc *BankAccount) Owner() string {
	return acc.owner
}

func (acc *BankAccount) Transfer(to *BankAccount, amount float64) error {
	if err := acc.Withdraw(amount); err != nil {
		return fmt.Errorf("ошибка перевода: %v", err)
	}
	if err := to.Deposit(amount); err != nil {
		acc.balance += amount
		return fmt.Errorf("ошибка зачисления: %v", err)
	}
	return nil
}

func main() {
	account1, err := NewBankAccount("Александр В", 1000)
	if err != nil {
		fmt.Println("Ошибка создания счета:", err)
		return
	}

	account2, err := NewBankAccount("Андрей П", 500)
	if err != nil {
		fmt.Println("Ошибка создания счета:", err)
		return
	}

	fmt.Printf("Счет %s: %.2f руб.\n", account1.Owner(), account1.Balance())
	fmt.Printf("Счет %s: %.2f руб.\n", account2.Owner(), account2.Balance())

	if err := account1.Deposit(300); err != nil {
		fmt.Println("Ошибка пополнения:", err)
	} else {
		fmt.Println("\nПосле пополнения:")
		fmt.Printf("Счет %s: %.2f руб.\n", account1.Owner(), account1.Balance())
	}

	if err := account2.Withdraw(200); err != nil {
		fmt.Println("Ошибка снятия:", err)
	} else {
		fmt.Println("\nПосле снятия:")
		fmt.Printf("Счет %s: %.2f руб.\n", account2.Owner(), account2.Balance())
	}

	fmt.Println("\nПеревод 400 руб. от Александра к Андрею")
	if err := account1.Transfer(account2, 400); err != nil {
		fmt.Println("Ошибка перевода:", err)
	} else {
		fmt.Println("После перевода:")
		fmt.Printf("Счет %s: %.2f руб.\n", account1.Owner(), account1.Balance())
		fmt.Printf("Счет %s: %.2f руб.\n", account2.Owner(), account2.Balance())
	}

	if err := account2.Withdraw(1000); err != nil {
		fmt.Println("\nПопытка снятия:", err)
	}
}
