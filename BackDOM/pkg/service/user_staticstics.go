package service

import (
	todo "BackDOM"
	"BackDOM/pkg/repository"
)

type UserStatisticsService struct {
	repo repository.UserStatistics
}

func NewUserStatisticsService(repo repository.UserStatistics) *UserStatisticsService {
	return &UserStatisticsService{repo: repo}
}

func (s *UserStatisticsService) GetUserStatisticsById(userId int) (todo.UserStatsRequest, error) {
	var UserStatsReq todo.UserStatsRequest
	UserStatistics, err := s.repo.GetUserStatisticsById(userId)
	if err != nil {
		return UserStatsReq, err
	}

	UserStatsReq.UserID = UserStatistics.UserID
	UserStatsReq.MoreDrinks = maxOfThree(UserStatistics.LightDrinks, UserStatistics.MediumDrinks, UserStatistics.HeavyDrinks)
	UserStatsReq.LastConsumptionDate = UserStatistics.LastConsumptionDate
	return UserStatsReq, err
}

func (s *UserStatisticsService) UpdateLastConsumptionDate(userId int, lastConsumptionDate string) error {
	return s.repo.UpdateLastConsumptionDate(userId, lastConsumptionDate)
}

func (s *UserStatisticsService) UpdateNumberOfDrinks(userId int, alcoholPercentage float64) error {
	var err error
	if alcoholPercentage < 10 {
		err = s.repo.UpdateLightDrinksStatistics(userId)
	} else if alcoholPercentage >= 10 && alcoholPercentage <= 25 {
		err = s.repo.UpdateMediumDrinksStatistics(userId)
	} else {
		err = s.repo.UpdateHeavyDrinksStatistics(userId)
	}
	return err
}

func maxOfThree(a, b, c float64) string {
	if a >= b && a >= c {
		return "Лёгкие напитки (до 10%)"
	} else if b >= a && b >= c {
		return "Средние напитки (от 10% до 25%)"
	} else {
		return "Тяжёлые напитки (от 25%)"
	}
}
