package service

import (
	"arquitetura-go/internal/email"
	"arquitetura-go/internal/products/domain"
	"context"
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

func (s service) GetAll(ctx context.Context) ([]domain.Product, error) {
	ps, err := s.repository.GetAll(ctx)
	if err != nil {
		return []domain.Product{}, err
	}

	return ps, nil
}

func (s service) Store(ctx context.Context, name, productType string, count int, price float64) (domain.Product, error) {
	product, err := s.repository.Store(ctx, 0, name, productType, count, price)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s service) Update(ctx context.Context, id int, name, productType string, count int, price float64) (domain.Product, error) {
	product, err := s.repository.Update(ctx, id, name, productType, count, price)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s service) UpdateName(ctx context.Context, id int, name string) (domain.Product, error) {
	product, err := s.repository.UpdateName(ctx, id, name)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
