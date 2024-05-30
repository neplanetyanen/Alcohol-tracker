package repository

import (
	todo "BackDOM"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserStatisticsPostgres struct {
	db *sqlx.DB
}

func NewUserStatisticsPostgres(db *sqlx.DB) *UserStatisticsPostgres {
	return &UserStatisticsPostgres{db: db}
}

func (r *UserStatisticsPostgres) GetUserStatisticsById(userId int) (todo.UserStatistics, error) {
	var userStatistics todo.UserStatistics

	query := fmt.Sprintf("SELECT user_id, light_drinks, medium_drinks, heavy_drinks, last_consumption_date FROM %s WHERE user_id = $1", userStatisticsTable)

	err := r.db.Get(&userStatistics, query, userId)
	if err != nil {
		return userStatistics, fmt.Errorf("failed to execute query: %v", err)
	}

	return userStatistics, nil
}

func (r *UserStatisticsPostgres) UpdateLastConsumptionDate(userId int, lastConsumptionDate string) error {
	query := fmt.Sprintf(`
        UPDATE %s
        SET last_consumption_date = $1
        WHERE user_id = $2
    `, userStatisticsTable)
	_, err := r.db.Exec(query, lastConsumptionDate, userId)
	return err
}

func (r *UserStatisticsPostgres) UpdateLightDrinksStatistics(userId int) error {
	var query string
	query = `UPDATE user_statistics SET light_drinks = light_drinks + 1 WHERE user_id = $1`
	_, err := r.db.Exec(query, userId)
	return err
}

func (r *UserStatisticsPostgres) UpdateMediumDrinksStatistics(userId int) error {
	var query string
	query = `UPDATE user_statistics SET medium_drinks = medium_drinks + 1 WHERE user_id = $1`
	_, err := r.db.Exec(query, userId)
	return err
}

func (r *UserStatisticsPostgres) UpdateHeavyDrinksStatistics(userId int) error {
	var query string
	query = `UPDATE user_statistics SET heavy_drinks = heavy_drinks + 1 WHERE user_id = $1`
	_, err := r.db.Exec(query, userId)
	return err
}
