package service

import (
	"github.com/rfaelmarini/call-cat-api/entity"
	"github.com/rfaelmarini/call-cat-api/repository"
)

type ResponseService interface {
	Save(entity.Response) entity.Response
	Find(url string) entity.Response
}

type responseService struct {
	lastResponse       entity.Response
	responseRepository repository.ResponseRepository
}

func NewResponseService(repo repository.ResponseRepository) ResponseService {
	return &responseService{
		responseRepository: repo,
	}
}

func (service *responseService) Save(response entity.Response) entity.Response {
	service.lastResponse = response
	service.responseRepository.Save(response)
	return response
}

func (service *responseService) Find(url string) entity.Response {
	if service.lastResponse.RequestedURL == url {
		return service.lastResponse
	}

	return service.responseRepository.Find(url)
}
