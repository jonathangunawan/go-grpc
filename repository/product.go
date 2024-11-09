package repository

import (
	"fmt"

	"github.com/jonathangunawan/go-grpc/constant"
	"github.com/jonathangunawan/go-grpc/entity"
)

type ProductRepoItf interface {
	InsertProduct(data entity.Product) (*entity.Product, error)
	GetProduct() []entity.Product
}

type ProductRepoImpl struct {
	counter int64
	db      []entity.ProductModel
	name    map[string]struct{}
}

func NewProductRepo() *ProductRepoImpl {
	return &ProductRepoImpl{
		counter: 0,
		db:      []entity.ProductModel{},
		name:    make(map[string]struct{}),
	}
}

func (pr *ProductRepoImpl) GetProduct() []entity.Product {
	var res []entity.Product

	for _, val := range pr.db {
		res = append(res, entity.Product(val))
	}

	return res
}

func (pr *ProductRepoImpl) InsertProduct(data entity.Product) (*entity.Product, error) {
	var res *entity.Product

	// if product name already added then return error
	_, ok := pr.name[data.Name]
	if ok {
		return nil, fmt.Errorf(constant.ErrUniqueName)
	}

	// add the data
	// increase the counter for id
	// then the data to db slice
	// then to the checker
	pr.counter++
	pr.db = append(pr.db, entity.ProductModel{
		ID:          pr.counter,
		Name:        data.Name,
		Description: data.Description,
	})
	pr.name[data.Name] = struct{}{}

	// assigning the data to entity
	res = &data
	res.ID = pr.counter

	return res, nil
}
