package service

import "database/sql"

type Transaction struct {
	ID     int    `json:"id"`
	Memo   string `json:"memo"`
	Debit  string `json:"debit"`
	Credit int    `json:"credit"`
}

type TransactionService struct {
	db *sql.DB
}

func NewTransactionService(db *sql.DB) *TransactionService {
	return &TransactionService{db: db}
}

func (s *TransactionService) GetTransactions() ([]Transaction, error) {
	query := "SELECT id, memo, debit, credit FROM transactions"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		err := rows.Scan(&transaction.ID, &transaction.Memo, &transaction.Debit, &transaction.Credit)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (s *TransactionService) GetTransactionByID(id int) (*Transaction, error) {
	query := "SELECT id, memo, debit, credit FROM transactions WHERE id = ?"
	row, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	var transaction Transaction
	err = row.Scan(&transaction.ID, &transaction.Memo, &transaction.Debit, &transaction.Credit)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (s *TransactionService) CreateTransaction(transaction *Transaction) error {
	query := "INSERT INTO transactions (id, memo, debit, credit) VALUES (?, ?, ?, ?)"
	result, err := s.db.Exec(query, transaction.ID, transaction.Memo, transaction.Debit, transaction.Credit)
	if err != nil {
		return err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	transaction.ID = int(lastInsertID)
	return nil
}

func (s *TransactionService) UpdateTransaction(transaction *Transaction) error {
	query := "UPDATE transactions SET memo = ?, debit = ?, credit = ? WHERE id = ?"
	_, err := s.db.Query(query, transaction.Memo, transaction.Debit, transaction.Credit)
	return err
}

func (s *TransactionService) DeleteTransaction(id int) error {
	query := "DELETE transactions WHERE id = ?"
	_, err := s.db.Query(query, id)
	return err
}
