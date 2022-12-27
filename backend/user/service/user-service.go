package service

import (
	"errors"

	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"github.com/kleberalves/problemCompanyApp/backend/services"
	"github.com/kleberalves/problemCompanyApp/backend/user"
)

type service struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) user.Service {
	return &service{
		repo: repo,
	}
}

func (srv *service) FindAll() (res []schema.UserRead, err error) {
	return srv.repo.FindAll()
}

func (srv *service) Create(input schema.User) (schema.User, error) {

	hash, _ := services.HashPassword(input.Password)

	user := schema.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hash,
		Profiles:  input.Profiles}

	return srv.repo.Create(user)
}

func (srv *service) Get(id int) (schema.UserRead, error) {
	return srv.repo.Get(id)
}

func (srv *service) Update(input schema.User) error {

	if input.ID <= 0 {
		return errors.New("ID not defined")
	}

	if input.Password != "" {
		//TODO: check password rules
		hash, _ := services.HashPassword(input.Password)
		input.Password = hash
	}

	return srv.repo.Update(input)
}

func (srv *service) Delete(ids []int) error {
	return srv.repo.Delete(ids)
}
