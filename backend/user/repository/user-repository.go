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
		Preload("Profiles", func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("User")
		}).
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
		Preload("Profiles",
			func(tx *gorm.DB) *gorm.DB {
				return tx.Select("type, user_id")
			}).
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

func (repo *repository) GetByEmail(email string) (schema.User, error) {
	var user schema.User
	err := repo.Conn.Model(&schema.User{}).
		Preload("Profiles",
			func(tx *gorm.DB) *gorm.DB {
				return tx.Omit("User")
			}).
		Where("email = ?", email).
		First(&user).Error
	return user, err
}
