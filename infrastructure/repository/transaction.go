package TransactionRepository

import "database/sql"

type TransactionRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db: db}	
}

func (t *TransactionRepositoryDb) SaveTransaction(transacion domain.Transaction, creditCard domain.CreditCard) error {
	stmt, err := t.db.Prepare( query:`insert into transactions(id, credit_card_id, amount, status, description, store, created_at) 
								values($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		transacion.ID, 
		transaction.CreditCardID, 
		transaction.Amount, 
		transacion.Status, 
		transacion.Description, 
		transaction.Store, 
		transaction.CreatedAt,
	)

	if err != nil {
		return err
	}

	if transacion.Status == "approved" {
		
	}
	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

