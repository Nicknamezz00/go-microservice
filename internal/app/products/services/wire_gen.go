// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package services

import (
	"github.com/Nicknamezz00/go-microservice/api/proto"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/config"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/log"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateProductsService(f string, detailsSvc proto.DetailsClient, ratingsSvc proto.RatingsClient, reviewsSvc proto.ReviewsClient) (ProductsService, error) {
	viper, err := config.NewViper(f)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	productsService := NewProductService(logger, detailsSvc, ratingsSvc, reviewsSvc)
	return productsService, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, ProviderSet)
