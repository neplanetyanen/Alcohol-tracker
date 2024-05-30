package repository

import (
	todo "BackDOM"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

const (
	initialNumberOfDrinks = 0
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int

	allergiesStr := strings.Join(user.Allergies, ",") // Преобразование среза в строку

	var query = fmt.Sprintf("INSERT INTO %s (username, email, password, nationality, height, weight, body_type, allergies, goal) values ($1, $2, $3, $4, $5, $6, $7, '%s', $8) RETURNING id", usersTable, allergiesStr)

	row := r.db.QueryRow(query, append([]interface{}{user.Username, user.Email, user.Password, user.Nationality, user.Height, user.Weight, user.BodyType}, user.Goal)...)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}

func (r *AuthPostgres) CreateUserStatistics(userId int) error {
	var id int
	var initialLastConsumptionDate time.Time

	query := fmt.Sprintf(`
        INSERT INTO %s (user_id, light_drinks, medium_drinks, heavy_drinks, last_consumption_date)
        VALUES ($1, $2, $3, $4, $5) RETURNING user_id`, userStatisticsTable)

	err := r.db.QueryRow(query, userId, initialNumberOfDrinks, initialNumberOfDrinks, initialNumberOfDrinks, initialLastConsumptionDate).Scan(&id)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}
