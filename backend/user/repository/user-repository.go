package repository

import (
	"github.com/kleberalves/problemCompanyApp/backend/enums"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"github.com/kleberalves/problemCompanyApp/backend/user"
	"gorm.io/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) user.Repository {
	return &userRepository{Conn}
}

func (repo *userRepository) AddProfile(userId int, typ enums.TypeUser) (schema.Profile, error) {

	profile := schema.Profile{
		UserID: userId,
		Type:   typ.EnumIndex(),
	}

	err := repo.Conn.Create(&profile).Error

	return profile, err

}
func (repo *userRepository) DeleteProfile(userId int, typ enums.TypeUser) error {
	profileWhere := schema.Profile{
		UserID: userId,
		Type:   typ.EnumIndex(),
	}

	//Unscoped() deletes permanently
	err := repo.Conn.Unscoped().Delete(&schema.Profile{}, repo.Conn.Where(profileWhere)).Error
	return err

}

func (repo *userRepository) FindAll() (res []schema.UserRead, err error) {

	var users []schema.UserRead
	errExec := repo.Conn.Model(&schema.User{}).
		Preload("Profiles").
		Find(&users).Error

	if errExec != nil {
		panic("Failed to retrieve all Users: " + errExec.Error())
	}

	return users, errExec
}

func (repo *userRepository) Create(input schema.User) (schema.User, error) {

	err := repo.Conn.Create(&input).Error
	return input, err

}

func (repo *userRepository) Get(id int) (schema.UserRead, error) {

	var user schema.UserRead
	err := repo.Conn.Model(&schema.User{}).
		Preload("Profiles").
		First(&user, id).Error
	return user, err

}

func (repo *userRepository) Update(user schema.User) error {

	err := repo.Conn.Model(&user).Updates(&user).Error
	return err

}

func (repo *userRepository) Delete(ids []int) error {

	//Soft deletes
	err := repo.Conn.Delete(&schema.User{}, ids).Error
	return err

}
