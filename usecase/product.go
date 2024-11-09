package usecase

import (
	"github.com/jonathangunawan/go-grpc/entity"
	"github.com/jonathangunawan/go-grpc/repository"
)

type ProductUsecaseItf interface {
	AddProduct(entity.Product) (*entity.Product, error)
	FindProduct() []entity.Product
}

type ProductUsecaseImpl struct {
	productRepo repository.ProductRepoItf
}

func NewProductUsecase(pr repository.ProductRepoItf) ProductUsecaseImpl {
	return ProductUsecaseImpl{
		productRepo: pr,
	}
}

func (puc ProductUsecaseImpl) FindProduct() []entity.Product {
	res := puc.productRepo.GetProduct()

	return res
}

func (puc ProductUsecaseImpl) AddProduct(data entity.Product) (*entity.Product, error) {
	res, err := puc.productRepo.InsertProduct(data)
	if err != nil {
		return nil, err
	}

	return res, nil
}
