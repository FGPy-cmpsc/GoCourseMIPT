package usecase

type TransferUseCase struct {
	dbManager DBAccountManager
}

func NewTransferUseCase(dbManager DBAccountManager) *TransferUseCase {
	return &TransferUseCase{dbManager: dbManager}
}

type TransferInput struct {
	SenderID   int64
	ReceiverID int64
	Amount     int64
}

func (transferUC *TransferUseCase) Execute(input TransferInput) error {
	sender, err := transferUC.dbManager.GetBankAccount(input.SenderID)
	if err != nil {
		return err
	}
	receiver, err := transferUC.dbManager.GetBankAccount(input.ReceiverID)
	if err != nil {
		return err
	}
	err = sender.TransferTo(receiver, input.Amount)
	if err != nil {
		return err
	}
	err = transferUC.dbManager.SaveBankAccount(sender)
	if err != nil {
		return err
	}
	err = transferUC.dbManager.SaveBankAccount(receiver)
	if err != nil {
		return err
	}
	return nil
}
