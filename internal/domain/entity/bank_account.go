package entity

import "errors"

type BankAccount struct {
	ID      int64
	Balance int64
}

func (bankAccount *BankAccount) GetBalance() int64 {
	return bankAccount.Balance
}

func (bankAccount *BankAccount) TransferTo(receiverBackAccount *BankAccount, amount int64) error {
	if amount <= 0 {
		return errors.New("forbidden non-positive amount to transfer")
	}
	if bankAccount.Balance < amount {
		return errors.New("insufficient funds")
	}
	bankAccount.Balance -= amount
	receiverBackAccount.Balance += amount
	return nil
}

func (bankAccount *BankAccount) Deposit(amount int64) error {
	if amount <= 0 {
		return errors.New("forbidden non-positive amount to deposit")
	}
	bankAccount.Balance += amount
	return nil
}
