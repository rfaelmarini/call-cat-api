package service

import "github.com/rfaelmarini/call-cat-api/entity"

type ResponseService interface {
	Save(entity.Response) entity.Response
	Find(url string) entity.Response
}

type responseService struct {
	lastResponse entity.Response
	responses    []entity.Response
}

func New() ResponseService {
	responseService := responseService{}
	return &responseService
}

func (service *responseService) Save(response entity.Response) entity.Response {
	service.lastResponse = response
	service.responses = append(service.responses, response)
	return response
}

func (service *responseService) Find(url string) entity.Response {
	if service.lastResponse.RequestedURL == url {
		return service.lastResponse
	}

	findedResponse := entity.Response{}
	for _, response := range service.responses {
		if response.RequestedURL == url {
			findedResponse = response
			break
		}
	}

	return findedResponse
}
