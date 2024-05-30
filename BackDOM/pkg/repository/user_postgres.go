package repository

import (
	todo "BackDOM"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetById(userId int) (todo.UserData, error) {
	var user todo.UserData

	query := `SELECT id, username, email, nationality, height, weight, body_type, allergies, goal FROM users WHERE id = $1`
	err := r.db.Get(&user, query, userId)
	if err != nil {
		return user, err
	}

	user.Allergies = strings.Split(user.Allergy, ",") // Разделение строки на элементы списка

	return user, nil
}

func (r *UserPostgres) GetId(username string) (int, error) {
	var userId int
	query := `
        SELECT id
        FROM users
        WHERE username = $1
        LIMIT 1;
    `
	err := r.db.QueryRow(query, username).Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return userId, nil
}

func (r *UserPostgres) Update(user todo.UserResponse) error {
	query := `
        UPDATE users
        SET username = $1, email = $2, nationality = $3, height = $4, weight = $5, body_type = $6, allergies = $7, goal = $8
        WHERE id = $9
    `
	_, err := r.db.Exec(query, user.Username, user.Email, user.Nationality, user.Height, user.Weight, user.BodyType, strings.Join(user.Allergies, ","), user.Goal, user.ID)
	return err
}

func (r *UserPostgres) UpdateTrackerScore(userId int, trackerScore time.Time) error {
	query := `UPDATE users SET tracker_score = $1 WHERE id = $2`
	_, err := r.db.Exec(query, trackerScore, userId)
	return err
}

func (r *UserPostgres) GetTrackerScoreById(userId int) (time.Time, error) {
	var trackerScore time.Time
	query := `SELECT tracker_score FROM users WHERE id=$1`
	row := r.db.QueryRow(query, userId)
	err := row.Scan(&trackerScore)
	if err != nil {
		return trackerScore, err
	}
	return trackerScore, nil
}
