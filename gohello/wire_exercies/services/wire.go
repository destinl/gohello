// wire.go
//go:build wireinject
// +build wireinject

package services

import "github.com/google/wire"

func InitializeService() (*Service, error) {
	// wire.Build(NewService, NewMySQLDatabase)
	// return nil, nil
	return wire.Build(NewService, NewMySQLDatabase).GetZeroValueAndError()
}

func NewService(db Database) *Service {
	return &Service{
		DB: db,
	}
}

func NewMySQLDatabase() *MySQLDatabase {
	return &MySQLDatabase{}
}
