package service

import (
	"arquitetura-go/internal/email"
	"arquitetura-go/internal/products/domain"
)

type service struct {
	repository domain.ProductRepository
	// email      email.ServiceEmail
}

func NewService(r domain.ProductRepository, e email.ServiceEmail) domain.ProductService {
	return &service{
		repository: r,
		// email:      e,
	}
}

func (s service) GetAll() ([]domain.Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return []domain.Product{}, err
	}

	return ps, nil
}

func (s service) Store(name, typee string, count int, price float64) (domain.Product, error) {
	lastID, err := s.repository.LastID()

	if err != nil {
		return domain.Product{}, err
	}

	lastID++

	product, err := s.repository.Store(lastID, name, typee, count, price)

	if err != nil {
		return domain.Product{}, err
	}

	// s.email.SendEmail(name)

	return product, nil

}

func (s service) Update(id int, name, productType string, count int, price float64) (domain.Product, error) {
	product, err := s.repository.Update(id, name, productType, count, price)
	if err != nil {
		return domain.Product{}, err
	}
	return product, err
}

func (s service) UpdateName(id int, name string) (domain.Product, error) {
	product, err := s.repository.UpdateName(id, name)
	if err != nil {
		return domain.Product{}, err
	}
	return product, err
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
