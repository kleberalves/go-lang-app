package service

import (
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
	return srv.repo.AddProfile(profileId, enums.TypeUser(typo))

}
func (srv *service) RemoveProfiles(userIds []int, typo int) error {
	return srv.repo.RemoveProfiles(userIds, enums.TypeUser(typo))
}

func (srv *service) FindAll() ([]schema.Profile, error) {
	return srv.repo.FindAll()
}
