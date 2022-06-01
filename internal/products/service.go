package products

import "arquitetura-go/internal/email"

type Service interface {
	GetAll() ([]Product, error)
	Store(name, typee string, count int, price float64) (Product, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
	email      email.ServiceEmail
}

func NewService(r Repository, e email.ServiceEmail) Service {
	return &service{
		repository: r,
		email:      e,
	}
}

func (s service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
	// return s.repository.GetAll()
}

func (s service) Store(name, typee string, count int, price float64) (Product, error) {
	lastID, err := s.repository.LastID()

	if err != nil {
		return Product{}, err
	}

	lastID++

	product, err := s.repository.Store(lastID, name, typee, count, price)

	if err != nil {
		return Product{}, err
	}

	s.email.SendEmail(name)

	return product, nil

}

func (s service) Update(id int, name, productType string, count int, price float64) (Product, error) {
	product, err := s.repository.Update(id, name, productType, count, price)
	if err != nil {
		return Product{}, err
	}
	return product, err
}

func (s service) UpdateName(id int, name string) (Product, error) {
	product, err := s.repository.UpdateName(id, name)
	if err != nil {
		return Product{}, err
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
