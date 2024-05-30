package service

import (
	todo "BackDOM"
	"BackDOM/pkg/repository"
)

type DrinkRecordService struct {
	repo repository.DrinkRecords
}

func NewDrinkRecordService(repo repository.DrinkRecords) *DrinkRecordService {
	return &DrinkRecordService{repo: repo}
}

func (s *DrinkRecordService) CreateDrinkRecord(record todo.DrinkRecord) error {
	return s.repo.CreateDrinkRecord(record)
}
