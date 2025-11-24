package db

import (
	"bank/internal/domain/entity"
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type FileDB struct {
	file *os.File
}

func NewFileDB(file *os.File) *FileDB {
	return &FileDB{file: file}
}

func (fileDB *FileDB) GetBankAccount(ID int64) (*entity.BankAccount, error) {
	if _, err := fileDB.file.Seek(0, 0); err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(fileDB.file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		curID, err := strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			return nil, err
		}
		balance, err := strconv.ParseInt(split[1], 10, 64)
		if err != nil {
			return nil, err
		}
		if curID == ID {
			return &entity.BankAccount{ID: curID, Balance: balance}, nil
		}
	}
	return nil, errors.New("no bank account with this ID")
}
func (fileDB *FileDB) SaveBankAccount(bankAccount *entity.BankAccount) error {
	if _, err := fileDB.file.Seek(0, 0); err != nil {
		return err
	}
	var lines []string
	scanner := bufio.NewScanner(fileDB.file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		curID, err := strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			return err
		}
		if curID == bankAccount.ID {
			lines = append(lines, strconv.FormatInt(curID, 10)+" "+strconv.FormatInt(bankAccount.Balance, 10))
		} else {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	err := os.WriteFile(fileDB.file.Name(), []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		panic(err)
	}
	return nil
}

func (fileDB *FileDB) SaveBankAccountsTransaction(bankAccounts []*entity.BankAccount) error {
	if _, err := fileDB.file.Seek(0, 0); err != nil {
		return err
	}
	var lines []string
	scanner := bufio.NewScanner(fileDB.file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		curID, err := strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			return err
		}
		foundID := false
		for _, bankAccount := range bankAccounts {
			if curID == bankAccount.ID {
				lines = append(lines, strconv.FormatInt(curID, 10)+" "+strconv.FormatInt(bankAccount.Balance, 10))
				foundID = true
				break
			}
		}
		if !foundID {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	err := os.WriteFile(fileDB.file.Name(), []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		panic(err)
	}
	return nil
}
