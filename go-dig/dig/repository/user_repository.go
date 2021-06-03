package repository

type IUserRepository interface {
	GetUserByID(id int)
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

type UserRepository struct{}

func (_ *UserRepository) GetUserByID(id int) {
	// RDBMS なり nosql なり
	return
}
