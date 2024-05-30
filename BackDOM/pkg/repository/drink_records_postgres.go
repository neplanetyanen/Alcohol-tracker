package repository

import (
	todo "BackDOM"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DrinkRecordPostgres struct {
	db *sqlx.DB
}

func NewDrinkRecordPostgres(db *sqlx.DB) *DrinkRecordPostgres {
	return &DrinkRecordPostgres{db: db}
}

func (r *DrinkRecordPostgres) CreateDrinkRecord(drinkRecord todo.DrinkRecord) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	var id int
	createDrinkRecordQuery := fmt.Sprintf("INSERT INTO %s (user_id, date, time, quantity, drink_name ) VALUES ($1, $2, $3, $4, $5) RETURNING id", drinkRecordsTable) // Используйте реальное имя таблицы
	row := tx.QueryRow(createDrinkRecordQuery, drinkRecord.UserId, drinkRecord.Date, drinkRecord.Time, drinkRecord.Quantity, drinkRecord.DrinkName)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to execute query: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
