package products

import "fmt"

//Repositorio

var ps []Product = []Product{}

var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name, typee string, count int, price float64) (Product, error)
	LastID() (int, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type repository struct{}

func (repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (repository) LastID() (int, error) {
	return lastID, nil
}

func (repository) Store(id int, name, typee string, count int, price float64) (Product, error) {
	p := Product{id, name, typee, count, price}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}

func (repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
	p := Product{Name: name, Type: productType, Count: count, Price: price}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("produto %d não encontrado", id)
	}
	return p, nil
}

func (repository) UpdateName(id int, name string) (Product, error) {
	var p Product
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Name = name
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("produto %d no encontrado", id)
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

	//Exemplo
	// 0 1 2 3 4 5
	//[1,2,3,4,5,6]
	//index = 2

	//ps[0:5] = ps[0:5] = [1,2,3,4,5]
	//ps[6:] = ps[3:] = [4,5,6]

	//[4,5,6]... = 4, 5 ,6
	//ps = append([1,2], [4,5,6]...)

	//[1,2,4,5,6]

	ps = append(ps[:index], ps[index+1:]...)
	return nil
}

func NewRepository() Repository {
	return &repository{}
}
