package repository

import (
	"github.com/kleberalves/problemCompanyApp/backend/enums"
	"github.com/kleberalves/problemCompanyApp/backend/profile"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"gorm.io/gorm"
)

type repository struct {
	Conn *gorm.DB
}

func NewProfileRepository(Conn *gorm.DB) profile.Repository {
	return &repository{Conn}
}

func (repo *repository) AddProfile(userId int, typ enums.TypeUser) (schema.Profile, error) {

	profile := schema.Profile{
		UserID: userId,
		Type:   typ.EnumIndex(),
	}

	err := repo.Conn.Model(&schema.Profile{}).Preload("User").Create(&profile).Error

	return profile, err

}
func (repo *repository) RemoveProfiles(userIds []int, typ enums.TypeUser) error {
	//Unscoped() deletes permanently
	err := repo.Conn.Unscoped().
		Delete(&schema.Profile{},
			repo.Conn.Where("user_id in ? AND type = ? ", userIds, typ.EnumIndex())).
		Error
	return err
}

func (repo *repository) FindAll() ([]schema.Profile, error) {

	var profiles []schema.Profile
	err := repo.Conn.Model(&schema.Profile{}).Preload("User", func(tx *gorm.DB) *gorm.DB {
		return tx.Omit("Password")
	}).Find(&profiles).Error

	return profiles, err
}
