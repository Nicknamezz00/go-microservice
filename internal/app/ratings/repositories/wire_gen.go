// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package repositories

import (
	"github.com/Nicknamezz00/go-microservice/internal/pkg/config"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/database"
	"github.com/Nicknamezz00/go-microservice/internal/pkg/log"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateRatingsRepository(f string) (RatingsRepository, error) {
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
	databaseOptions, err := database.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	db, err := database.NewDatabase(databaseOptions)
	if err != nil {
		return nil, err
	}
	ratingsRepository := NewMySQLRatingsRepository(logger, db)
	return ratingsRepository, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, ProviderSet)
