package service

import (
	todo "BackDOM"
	"BackDOM/pkg/repository"
	"time"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	CreateUserStatistics(userId int) error
}

type User interface {
	GetById(userId int) (todo.UserData, error)
	GetId(username string) (int, error)
	Update(user todo.UserResponse) (todo.UserData, error)
	CalculateUserNorm(userId int, alcoholPercentage float64) (float64, error)
	UpdateUserTrackerScore(userId int, size, alcoholPercentage float64) error
	GetUserTrackerScore(userId int) (time.Time, error)
}

type UserStatistics interface {
	GetUserStatisticsById(userId int) (todo.UserStatsRequest, error)
	UpdateLastConsumptionDate(userId int, lastConsumptionDate string) error
	UpdateNumberOfDrinks(userId int, alcoholPercentage float64) error
}

type DrinkRecords interface {
	CreateDrinkRecord(drinkRecord todo.DrinkRecord) error
}

type Service struct {
	Authorization
	User
	UserStatistics
	DrinkRecords
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:  NewAuthService(repos.Authorization),
		User:           NewUserService(repos.User),
		UserStatistics: NewUserStatisticsService(repos.UserStatistics),
		DrinkRecords:   NewDrinkRecordService(repos.DrinkRecords),
	}
}
