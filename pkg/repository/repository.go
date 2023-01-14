package repository

type ManipulateUser interface {
}

type Repository struct {
	ManipulateUser
}

func NewRepository() *Repository {
	return &Repository{}
}
