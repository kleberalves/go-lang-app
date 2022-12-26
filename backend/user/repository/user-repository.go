package repository

import (
	"fmt"

	"github.com/kleberalves/problemCompanyApp/backend/enums"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"github.com/kleberalves/problemCompanyApp/backend/services"
	"github.com/kleberalves/problemCompanyApp/backend/user"
	"gorm.io/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) user.Repository {
	return &userRepository{Conn}
}

func (repo *userRepository) AssociateProfile(userId int, typ enums.TypeUser) schema.Profile {

	profile := schema.Profile{
		UserID: userId,
		Type:   typ,
	}

	err := repo.Conn.Create(&profile).Error

	if err != nil {
		panic("Failed to associate profile to user: " + err.Error())
	}

	return profile

}
func (repo *userRepository) RemoveProfile(userId int, typ enums.TypeUser) {
	profile := schema.Profile{
		UserID: userId,
		Type:   typ,
	}

	err := repo.Conn.Delete(&profile, repo.Conn.Where(&schema.Profile{UserID: userId, Type: enums.TypeUser(typ.EnumIndex())})).Error

	if err != nil {
		panic("Failed to remove profile to user: " + err.Error())
	}

}

func (repo *userRepository) FindAll() (res []schema.UserRead, err error) {

	var profiles []schema.Profile
	errE := repo.Conn.Model(&schema.Profile{}).Find(&profiles).Error
	if errE != nil {
		panic("Failed to retrieve all profiles: " + err.Error())
	}

	fmt.Println(profiles)

	var users []schema.UserRead
	errExec := repo.Conn.Model(&schema.User{}).Preload("Profiles").Find(&users).Error

	if errExec != nil {
		panic("Failed to retrieve all Users: " + err.Error())
	}

	return users, errExec
}

func (repo *userRepository) Create(input schema.User) (schema.User, error) {

	hash, _ := services.HashPassword(input.Password)

	user := schema.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hash,
		Profiles:  input.Profiles}

	err := repo.Conn.Create(&user).Error

	return user, err
}
