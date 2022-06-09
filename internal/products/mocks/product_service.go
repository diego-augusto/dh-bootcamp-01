package mocks

import (
	"arquitetura-go/internal/products"

	"github.com/stretchr/testify/mock"
)

type ProductService struct {
	mock.Mock
}

func (p *ProductService) GetAll() ([]products.Product, error) {
	args := p.Called()

	var product []products.Product

	if rf, ok := args.Get(0).(func() []products.Product); ok {
		product = rf()
	} else {
		if args.Get(0) != nil {
			product = args.Get(0).([]products.Product)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return product, err
}

func (p *ProductService) Store(
	id int,
	name, typee string,
	count int,
	price float64,
) (products.Product, error) {
	args := p.Called()

	var product products.Product

	if rf, ok := args.Get(0).(func(
		id int,
		name, typee string,
		count int,
		price float64,
	) products.Product); ok {
		product = rf(id, name, typee, count, price)
	} else {
		if args.Get(0) != nil {
			product = args.Get(0).(products.Product)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return product, err
}

func (p *ProductService) Update(
	id int,
	name, productType string,
	count int,
	price float64,
) (products.Product, error) {
	args := p.Called()

	var product products.Product

	if rf, ok := args.Get(0).(func(
		id int,
		name, productType string,
		count int,
		price float64,
	) products.Product); ok {
		product = rf(id, name, productType, count, price)
	} else {
		if args.Get(0) != nil {
			product = args.Get(0).(products.Product)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return product, err
}

func (p *ProductService) UpdateName(
	id int,
	name string,
) (products.Product, error) {
	args := p.Called()

	var product products.Product

	if rf, ok := args.Get(0).(func(id int, name string) products.Product); ok {
		product = rf(id, name)
	} else {
		if args.Get(0) != nil {
			product = args.Get(0).(products.Product)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return product, err
}

func (p *ProductService) Delete(id int) error {
	args := p.Called()

	var err error

	if rf, ok := args.Get(0).(func(id int) error); ok {
		err = rf(id)
	} else {
		err = args.Error(0)
	}

	return err
}
