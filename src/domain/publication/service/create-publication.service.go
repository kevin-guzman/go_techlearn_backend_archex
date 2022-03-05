package service

import (
	"golang-gingonic-hex-architecture/src/domain/errors"
	"golang-gingonic-hex-architecture/src/domain/publication/model"
	"golang-gingonic-hex-architecture/src/domain/publication/port/repository"
	"net/http"
)

var (
	errTrace       string = "This error has ocurred in create-publication.service.go"
	internalError  string = "Internal server error"
	successMessage string = "Publication has succesfully created!"
)

type ServiceCreatePublication struct {
	publicationRepository repository.RepositoryPublication
}

func NewServiceCreatePublication(PublicationR repository.RepositoryPublication) *ServiceCreatePublication {
	return &ServiceCreatePublication{
		publicationRepository: PublicationR,
	}
}

func (sca *ServiceCreatePublication) Run(publication model.Publication) (string, error, int) {
	err := sca.publicationRepository.Save(publication)
	if err != nil {
		return "", errors.NewErrorCore(err, errTrace, "Service error").PublicError(), http.StatusInternalServerError
	}

	return successMessage, nil, http.StatusOK
}
