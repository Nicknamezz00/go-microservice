// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package services

import (
	"github.com/Nicknamezz00/go-microservice/internal/app/reviews/repositories"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/config"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/database"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/log"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateReviewsService(f string, r repositories.ReviewsRepository) (ReviewsService, error) {
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
	reviewsService := NewReviewsService(logger, r)
	return reviewsService, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, ProviderSet)
