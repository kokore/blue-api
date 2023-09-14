package product

import (
	"blue-api/internal/config"
	"blue-api/internal/repository"
	productRepo "blue-api/internal/repository/product"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type serviceImpl struct {
	appConfig         *config.AppConfig
	productRepository productRepo.Repo
}

type Service interface {
	CreateProductService(ctx context.Context, name string, price uint, currentstock uint, image string) error
	GetProductsService(ctx context.Context) ([]productRepo.Product, error)
	UpdateProductServie(ctx context.Context, id string, name string, price uint, currentstock uint, image string) error
}

func InitProductService(appConfig *config.AppConfig, repos repository.Repositories) Service {
	return &serviceImpl{
		appConfig:         appConfig,
		productRepository: repos.Product,
	}
}

func (s serviceImpl) CreateProductService(ctx context.Context, name string, price uint, currentstock uint, image string) error {
	err := s.productRepository.InsertOne(ctx, productRepo.NewFilter().SetName(name).SetPrice(price).SetCurrentStock(currentstock).SetImage(image).SetCreatedAt().SetUpdatedAt())
	if err != nil {
		return err
	}
	return nil
}

func (s serviceImpl) GetProductsService(ctx context.Context) ([]productRepo.Product, error) {
	products, err := s.productRepository.Find(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s serviceImpl) UpdateProductServie(ctx context.Context, productId string, name string, price uint, currentstock uint, image string) error {
	objectID, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return err
	}
	_, errRepo := s.productRepository.FindAndUpdate(ctx, objectID, productRepo.NewUpdate().SetName(name).SetPrice(price).SetCurrentStock(currentstock).SetImage(image))
	if errRepo != nil {
		return err
	}

	return nil
}
