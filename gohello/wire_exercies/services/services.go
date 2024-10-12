// services.go
package services

type Database interface {
	Query() string
}

type Service struct {
	DB Database
}

func (s *Service) DoSomething() string {
	return s.DB.Query()
}
