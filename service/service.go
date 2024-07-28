package service

import (
	"database/sql"

	"github.com/TheGitExplorer/transactions/config"
	"github.com/TheGitExplorer/transactions/entity"
)

func AddTransaction(transaction entity.Transaction) error {
	query := "INSERT INTO transactions (id, amount, type, parent_id) VALUES (?, ?, ?, ?)"
	_, err := config.DB.Exec(query, transaction.ID, transaction.Amount, transaction.Type, transaction.ParentID)
	return err
}

func GetTransaction(id int64) (entity.Transaction, bool, error) {
	var transaction entity.Transaction
	query := "SELECT id, amount, type, parent_id FROM transactions WHERE id = ?"
	row := config.DB.QueryRow(query, id)
	err := row.Scan(&transaction.ID, &transaction.Amount, &transaction.Type, &transaction.ParentID)
	if err == sql.ErrNoRows {
		return transaction, false, nil
	} else if err != nil {
		return transaction, false, err
	}
	return transaction, true, nil
}

func GetTransactionsByType(transactionType string) ([]int64, error) {
	query := "SELECT id FROM transactions WHERE type = ?"
	rows, err := config.DB.Query(query, transactionType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func CalculateSum(id int64) (float64, error) {
	sum, err := calculateSum(id)
	return sum, err
}

func calculateSum(id int64) (float64, error) {
	var sum float64
	query := "SELECT id, amount FROM transactions WHERE id = ?"
	row := config.DB.QueryRow(query, id)
	var transaction entity.Transaction
	if err := row.Scan(&transaction.ID, &transaction.Amount); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	sum += transaction.Amount

	childQuery := "SELECT id FROM transactions WHERE parent_id = ?"
	rows, err := config.DB.Query(childQuery, id)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var childID int64
		if err := rows.Scan(&childID); err != nil {
			return 0, err
		}
		childSum, err := calculateSum(childID)
		if err != nil {
			return 0, err
		}
		sum += childSum
	}

	return sum, nil
}
