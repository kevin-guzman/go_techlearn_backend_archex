package service

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/domain/comment/model"
	"golang-gingonic-hex-architecture/src/domain/comment/port/repository"
	"golang-gingonic-hex-architecture/src/domain/errors"
	publicationRepository "golang-gingonic-hex-architecture/src/domain/publication/port/repository"
)

type ServiceCreateComment struct {
	commentRepository     repository.RepositoryComment
	publicationRepository publicationRepository.RepositoryPublication
}

func NewServiceCreateComment(CommentR repository.RepositoryComment, PublicationR publicationRepository.RepositoryPublication) *ServiceCreateComment {
	return &ServiceCreateComment{
		commentRepository:     CommentR,
		publicationRepository: PublicationR,
	}
}

func (scc *ServiceCreateComment) Run(comment model.Comment) interface{} {
	const successMessage = "El comentario se ha creado de forma satisfactoria!"
	existPublication, err := scc.publicationRepository.ExistById(comment.PublicationId)
	if err != nil {
		return errors.NewErrorPort(err)
	}
	if !existPublication {
		err := fmt.Errorf("La publicación con id %d no exste!", comment.PublicationId)
		return errors.NewErPublicationDoesntExist(err, err.Error())
	}

	err = scc.commentRepository.Save(comment)
	if err != nil {
		return errors.NewErrorPort(err)
	}

	return successMessage
}
