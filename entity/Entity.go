package entity

import (
    "database/sql"
    "errors"
    "log"
)

type Transaction struct {
    ID       int64   `json:"id"`
    Amount   float64 `json:"amount"`
    Type     string  `json:"type"`
    ParentID int64  `json:"parent_id,omitempty"`
}

var db *sql.DB

func SetDB(database *sql.DB) {
    db = database
}

func AddTransaction(transaction Transaction) error {
    query := "INSERT INTO transactions (id, amount, type, parent_id) VALUES (?, ?, ?, ?)"
    _, err := db.Exec(query, transaction.ID, transaction.Amount, transaction.Type, transaction.ParentID)
    if err != nil {
        log.Println("Error inserting transaction:", err)
        return err
    }
    return nil
}

func GetTransaction(id int64) (Transaction, bool, error) {
    var transaction Transaction
    query := "SELECT id, amount, type, parent_id FROM transactions WHERE id = ?"
    row := db.QueryRow(query, id)
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
    rows, err := db.Query(query, transactionType)
    if err != nil {
        log.Println("Error retrieving transactions by type:", err)
        return nil, err
    }
    defer rows.Close()

    var ids []int64
    for rows.Next() {
        var id int64
        if err := rows.Scan(&id); err != nil {
            log.Println("Error scanning transaction id:", err)
            return nil, err
        }
        ids = append(ids, id)
    }

    return ids, nil
}

func CalculateSum(id int64) (float64, error) {
    transaction, exists, err := GetTransaction(id)
    if err != nil {
        return 0, err
    }
    if !exists {
        return 0, errors.New("transaction not found")
    }

    sum := transaction.Amount
    childQuery := "SELECT id FROM transactions WHERE parent_id = ?"
    rows, err := db.Query(childQuery, id)
    if err != nil {
        log.Println("Error retrieving child transactions:", err)
        return 0, err
    }
    defer rows.Close()

    for rows.Next() {
        var childID int64
        if err := rows.Scan(&childID); err != nil {
            log.Println("Error scanning child transaction id:", err)
            return 0, err
        }
        childSum, err := CalculateSum(childID)
        if err != nil {
            return 0, err
        }
        sum += childSum
    }

    return sum, nil
}
