package service

import (
	"errors"

	"github.com/kleberalves/problemCompanyApp/backend/enums"
	"github.com/kleberalves/problemCompanyApp/backend/profile"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

type service struct {
	repo profile.Repository
}

func NewProfileService(repo profile.Repository) profile.Service {
	return &service{
		repo: repo,
	}
}

func (srv *service) AddProfile(profileId int, typo int) (schema.Profile, error) {

	e := enums.TypeUser(typo)

	if e.String() == "" {
		return schema.Profile{}, errors.New("invalid-profile")
	}

	return srv.repo.AddProfile(profileId, e)

}
func (srv *service) RemoveProfiles(userIds []int, typo int) error {
	return srv.repo.RemoveProfiles(userIds, enums.TypeUser(typo))
}

func (srv *service) FindAll() ([]schema.Profile, error) {
	return srv.repo.FindAll()
}
