package service

import (
	"counter-api/domain"
	"counter-api/repository"
)

type CounterService interface {
	GetNumber()(*domain.Number, error)
	IncrementNumber()(*domain.Number, error)
	DecrementNumber()(*domain.Number, error)
}

type counterService struct {
	repository repository.CounterRepository
}

func NewCounterService(repository repository.CounterRepository) CounterService {
	return &counterService{repository: repository}
}

func (c counterService) GetNumber() (*domain.Number, error) {
	return c.repository.GetNumber()
}

func (c counterService) IncrementNumber() (*domain.Number, error) {
	return c.repository.IncrementNumber()
}

func (c counterService) DecrementNumber() (*domain.Number, error) {
	return c.repository.DecrementNumber()
}
