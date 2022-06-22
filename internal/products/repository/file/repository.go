package repository

import (
	"arquitetura-go/internal/products/domain"
	"arquitetura-go/pkg/store"
	"fmt"
)

//Repositorio

var ps []domain.Product = []domain.Product{}

// repositório em memória
// type repository struct{}

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]domain.Product, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return []domain.Product{}, err
	}
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].ID, nil
}

func (r *repository) Store(id int, name, productType string, count int, price float64) (domain.Product, error) {
	var ps []domain.Product
	if err := r.db.Read(&ps); err != nil {
		return domain.Product{}, err
	}
	p := domain.Product{
		ID:    id,
		Name:  name,
		Type:  productType,
		Count: count,
		Price: price,
	}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (repository) Update(id int, name, productType string, count int, price float64) (domain.Product, error) {
	p := domain.Product{Name: name, Type: productType, Count: count, Price: price}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return domain.Product{}, fmt.Errorf("produto %d não encontrado", id)
	}
	return p, nil
}

func (repository) UpdateName(id int, name string) (domain.Product, error) {
	var p domain.Product
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Name = name
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return domain.Product{}, fmt.Errorf("produto %d no encontrado", id)
	}
	return p, nil
}

func (repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range ps {
		if ps[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("produto %d nao encontrado", id)
	}

	ps = append(ps[:index], ps[index+1:]...)
	return nil
}

func NewRepository(db store.Store) domain.ProductRepository {
	return &repository{
		db: db,
	}
}
