package service

import (
	"database/sql"
)

type AccountService struct {
	db *sql.DB
}

type Account struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Balance int    `json:"balance"`
}

func NewAccountService(db *sql.DB) *AccountService {
	return &AccountService{db: db}
}

func (s *AccountService) GetAccounts() ([]Account, error) {
	query := "SELECT id, name, type, balance FROM accounts"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var accounts []Account
	for rows.Next() {
		var account Account
		err := rows.Scan(&account.ID, &account.Name, &account.Type, &account.Balance)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (s *AccountService) GetAccountByID(id int) (*Account, error) {
	query := "SELECT id, name, type, balance FROM accounts WHERE id = ?"
	row, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	var account Account
	err = row.Scan(&account.ID, &account.Name, &account.Type, &account.Balance)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (s *AccountService) GetAccount(account *Account) error {
	query := "INSERT INTO accounts (id, name, type, balance) VALUES (?, ?, ?, ?)"
	result, err := s.db.Exec(query, account.ID, account.Name, account.Type, account.Balance)
	if err != nil {
		return err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	account.ID = int(lastInsertID)
	return nil
}

func (s *AccountService) CreateAccount(account *Account) error {
	query := "INSERT INTO accounts (id, name, type, balance) VALUES (?, ?, ?, ?)"
	result, err := s.db.Exec(query, account.ID, account.Name, account.Type, account.Balance)
	if err != nil {
		return err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	account.ID = int(lastInsertID)
	return nil
}

func (s *AccountService) UpdateAccount(account *Account) error {
	query := "UPDATE accounts SET name = ?, type = ? WHERE id = ?"
	_, err := s.db.Query(query, account.Name, account.Type, account.ID)
	return err
}

func (s *AccountService) DeleteAccount(id int) error {
	query := "DELETE accounts WHERE id = ?"
	_, err := s.db.Query(query, id)
	return err
}
