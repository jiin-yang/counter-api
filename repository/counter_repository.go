package repository

import "counter-api/domain"

type CounterRepository interface {
	GetNumber()(*domain.Number, error)
	IncrementNumber()(*domain.Number, error)
	DecrementNumber()(*domain.Number, error)
}

type counterRepository struct {
	number *domain.Number
}

func NewCounterRepository(number *domain.Number) CounterRepository {
	return &counterRepository{number: number}
}

func (c counterRepository) GetNumber() (*domain.Number, error) {
	return c.number, nil
}

func (c counterRepository) IncrementNumber() (*domain.Number, error) {
	c.number.Number++
	return c.number, nil
}

func (c counterRepository) DecrementNumber() (*domain.Number, error) {
	c.number.Number--
	return c.number, nil
}


