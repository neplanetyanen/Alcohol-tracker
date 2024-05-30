package repository

import (
	todo "BackDOM"
	"github.com/jmoiron/sqlx"
	"time"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
	CreateUserStatistics(userId int) error
}

type User interface {
	GetById(userId int) (todo.UserData, error)
	GetId(username string) (int, error)
	Update(response todo.UserResponse) error
	UpdateTrackerScore(userId int, trackerScore time.Time) error
	GetTrackerScoreById(userId int) (time.Time, error)
}

type UserStatistics interface {
	GetUserStatisticsById(userId int) (todo.UserStatistics, error)
	UpdateLastConsumptionDate(userId int, lastConsumptionDate string) error
	UpdateLightDrinksStatistics(userId int) error
	UpdateMediumDrinksStatistics(userId int) error
	UpdateHeavyDrinksStatistics(userId int) error
}

type DrinkRecords interface {
	CreateDrinkRecord(drinkRecord todo.DrinkRecord) error
}

type Repository struct {
	Authorization
	User
	UserStatistics
	DrinkRecords
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:  NewAuthPostgres(db),
		User:           NewUserPostgres(db),
		UserStatistics: NewUserStatisticsPostgres(db),
		DrinkRecords:   NewDrinkRecordPostgres(db),
	}
}
