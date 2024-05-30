package service

import (
	todo "BackDOM"
	"BackDOM/pkg/repository"
	"fmt"
	"time"
)

// Коэффициенты
var bodyTypeCoefficients = map[string]float64{
	"ectomorphic": 0.8,
	"mesomorphic": 1.2,
	"endomorphic": 1.0,
}

var nationalityCoefficients = map[string]float64{
	"north":     0.8,
	"northeast": 0.9,
	"east":      1.0,
	"southeast": 1.1,
	"south":     1.2,
	"southwest": 1.1,
	"west":      1.0,
	"northwest": 0.9,
}

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CalculateUserNorm(userId int, alcoholPercentage float64) (float64, error) {
	userData, err := s.repo.GetById(userId)
	if err != nil {
		return 0, err
	}

	bodyTypeCoeff, ok := bodyTypeCoefficients[userData.BodyType]
	if !ok {
		return 0, fmt.Errorf("неизвестный тип телосложения: %s", userData.BodyType)
	}

	nationalityCoeff, ok := nationalityCoefficients[userData.Nationality]
	if !ok {
		return 0, fmt.Errorf("неизвестная национальность: %s", userData.Nationality)
	}

	result := calculateNorm(bodyTypeCoeff, nationalityCoeff, userData.Weight, alcoholPercentage)
	return result, nil
}

func calculateNorm(bodyTypeCoeff, nationalityCoeff, weight, alcoholPercentage float64) float64 {
	return 1.2 * 65 * bodyTypeCoeff * nationalityCoeff * weight / alcoholPercentage
}

func (s *UserService) GetById(userId int) (todo.UserData, error) {
	return s.repo.GetById(userId)
}

func (s *UserService) GetId(username string) (int, error) {
	return s.repo.GetId(username)
}

func (s *UserService) Update(user todo.UserResponse) (todo.UserData, error) {
	err := s.repo.Update(user)
	if err != nil {
		return todo.UserData{}, err
	}
	return s.repo.GetById(user.ID)
}

func (s *UserService) UpdateUserTrackerScore(userId int, size, alcoholPercentage float64) error {
	trackerScore, err := s.GetUserTrackerScore(userId)
	if err != nil {
		return err
	}

	userData, err := s.repo.GetById(userId)
	if err != nil {
		return err
	}

	minutes := calculateTracker(size, userData.Weight, alcoholPercentage)
	duration := time.Minute * time.Duration(minutes)

	if trackerScore.Before(time.Now()) {
		trackerScore = time.Now()
	}

	trackerScore = trackerScore.Add(duration)
	return s.repo.UpdateTrackerScore(userId, trackerScore)
}

func calculateTracker(size float64, weight, alcoholPercentage float64) int64 {
	return int64((size * alcoholPercentage * 457) / (weight * 0.6 * 100))
}

func (s *UserService) GetUserTrackerScore(userId int) (time.Time, error) {
	return s.repo.GetTrackerScoreById(userId)
}
