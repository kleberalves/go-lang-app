package repository

import (
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"github.com/kleberalves/problemCompanyApp/backend/user"
	"gorm.io/gorm"
)

type repository struct {
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) user.Repository {
	return &repository{Conn}
}

func (repo *repository) FindAll() (res []schema.UserRead, err error) {

	var users []schema.UserRead
	errExec := repo.Conn.Model(&schema.User{}).
		Preload("Profiles").
		Find(&users).Error

	if errExec != nil {
		panic("Failed to retrieve all Users: " + errExec.Error())
	}

	return users, errExec
}

func (repo *repository) Create(input schema.User) (schema.User, error) {

	err := repo.Conn.Create(&input).Error
	return input, err

}

func (repo *repository) Get(id int) (schema.UserRead, error) {

	var user schema.UserRead
	err := repo.Conn.Model(&schema.User{}).
		Preload("Profiles").
		First(&user, id).Error
	return user, err

}

func (repo *repository) Update(user schema.User) error {

	err := repo.Conn.Model(&user).Updates(&user).Error
	return err

}

func (repo *repository) Delete(ids []int) error {

	//Soft deletes
	err := repo.Conn.Delete(&schema.User{}, ids).Error
	return err

}
