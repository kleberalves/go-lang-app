package service

import (
	"errors"

	"github.com/kleberalves/problemCompanyApp/backend/enums"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"github.com/kleberalves/problemCompanyApp/backend/services"
	"github.com/kleberalves/problemCompanyApp/backend/user"
)

type userService struct {
	userRepo user.Repository
}

func NewUserService(repo user.Repository) user.Service {
	return &userService{
		userRepo: repo,
	}
}

func (srv *userService) AddProfile(userId int, typo int) (schema.Profile, error) {
	return srv.userRepo.AddProfile(userId, enums.TypeUser(typo))

}
func (srv *userService) DeleteProfile(userId int, typo int) error {
	return srv.userRepo.DeleteProfile(userId, enums.TypeUser(typo))
}

func (srv *userService) FindAll() (res []schema.UserRead, err error) {
	return srv.userRepo.FindAll()
}

func (srv *userService) Create(input schema.User) (schema.User, error) {

	hash, _ := services.HashPassword(input.Password)

	user := schema.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hash,
		Profiles:  input.Profiles}

	return srv.userRepo.Create(user)
}

func (srv *userService) Get(id int) (schema.UserRead, error) {
	return srv.userRepo.Get(id)
}

func (srv *userService) Update(input schema.User) error {

	if input.ID <= 0 {
		return errors.New("ID not defined")
	}

	if input.Password != "" {
		//TODO: check password rules
		hash, _ := services.HashPassword(input.Password)
		input.Password = hash
	}

	return srv.userRepo.Update(input)
}

func (srv *userService) Delete(ids []int) error {
	return srv.userRepo.Delete(ids)
}
