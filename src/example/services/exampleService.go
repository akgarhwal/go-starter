package services

import (
	"go-starter/src/example/dto"
	"go-starter/src/example/models"
	"go-starter/src/example/repositories"
)

type ExampleService interface {
	GetExample(dto.ExampleRequestBody) models.Example
}

type exampleService struct {
	ExampleRepository repositories.ExampleRepositoryInterface
}

func NewExampleService(exampleRepository repositories.ExampleRepositoryInterface) ExampleService {
	return &exampleService{
		ExampleRepository: exampleRepository,
	}
}

func (es *exampleService) GetExample(request dto.ExampleRequestBody) models.Example {
	return es.ExampleRepository.Find(request.Name)
}
