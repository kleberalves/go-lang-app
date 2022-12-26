package service

import (
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

func (srv *userService) AssociateProfile(userId int, typo int) schema.Profile {
	return srv.userRepo.AssociateProfile(userId, enums.TypeUser(typo))

}
func (srv *userService) RemoveProfile(userId int, typo int) {
	srv.userRepo.RemoveProfile(userId, enums.TypeUser(typo))
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
