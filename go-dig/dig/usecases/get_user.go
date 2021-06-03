package usecases

import "go-dig/repository"

type GetUser interface {
}

func NewGetUser(userRepository repository.IUserRepository) GetUser {
	return &getUser{
		UserRepository : userRepository,
	}
}

type getUser struct {
	UserRepository repository.IUserRepository
}

