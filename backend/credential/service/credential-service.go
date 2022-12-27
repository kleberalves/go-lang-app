package service

import (
	"errors"

	"github.com/kleberalves/problemCompanyApp/backend/credential"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"github.com/kleberalves/problemCompanyApp/backend/services/security"
	"github.com/kleberalves/problemCompanyApp/backend/user"
)

type service struct {
	repo user.Repository
}

func NewCredentialService(repo user.Repository) credential.Service {
	return &service{
		repo: repo,
	}
}

func (srv *service) Login(user schema.UserCredential) (schema.Credential, error) {
	dbUser, err := srv.repo.GetByEmail(user.Email)

	if dbUser.Email == "" {
		err = errors.New("user-not-found")
	} else {
		if security.CheckPasswordHash(user.Password, dbUser.Password) {

			token, err := security.GenerateToken(dbUser.ID)

			if err == nil {
				credential := schema.Credential{
					FirstName: dbUser.FirstName,
					Email:     dbUser.Email,
					Profiles:  dbUser.Profiles,
					JwToken:   token,
				}
				return credential, err
			}
		} else {
			err = errors.New("invalid-password")
		}
	}

	return schema.Credential{}, err
}
