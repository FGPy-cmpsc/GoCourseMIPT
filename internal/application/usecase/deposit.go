package usecase

type DepositUseCase struct {
	dbManager DBAccountManager
}

func NewDepositUseCase(dbManager DBAccountManager) *DepositUseCase {
	return &DepositUseCase{dbManager: dbManager}
}

type DepositInput struct {
	ID     int64
	Amount int64
}

func (depositUC *DepositUseCase) Execute(input DepositInput) error {
	bankAccount, err := depositUC.dbManager.GetBankAccount(input.ID)
	if err != nil {
		return err
	}
	err = bankAccount.Deposit(input.Amount)
	if err != nil {
		return err
	}
	err = depositUC.dbManager.SaveBankAccount(bankAccount)
	return err
}
