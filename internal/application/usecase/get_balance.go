package usecase

import (
	"bank/internal/domain/entity"
	"fmt"
)

type DBAccountManager interface {
	GetBankAccount(ID int64) (*entity.BankAccount, error)
	SaveBankAccount(bankAccount *entity.BankAccount) error
}

type GetBalanceUseCase struct {
	dbManager DBAccountManager
}

func NewGetBalanceUseCase(dbManager DBAccountManager) *GetBalanceUseCase {
	return &GetBalanceUseCase{dbManager: dbManager}
}

type GetBalanceInput struct {
	ID int64
}

type GetBalanceOutput struct {
	Balance int64
}

func (useCase *GetBalanceUseCase) Execute(input GetBalanceInput) (*GetBalanceOutput, error) {
	fmt.Println("in GetBalance Execute")
	bankAccount, err := useCase.dbManager.GetBankAccount(input.ID)
	if err != nil {
		return nil, err
	}
	balance := bankAccount.GetBalance()
	return &GetBalanceOutput{Balance: balance}, nil
}
