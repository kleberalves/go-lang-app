package service

import (
	"errors"

	"github.com/kleberalves/problemCompanyApp/backend/credential"
	"github.com/kleberalves/problemCompanyApp/backend/enums"
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

			token, err := security.GenerateToken(int(dbUser.ID))

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

func (srv *service) CheckRolesByUserID(userId int, r []enums.TypeUser) (bool, error) {
	user, err := srv.repo.Get(userId)

	if err != nil {
		return false, err
	}

	var checkRoles bool

	for i := 0; i < len(r); i++ {
		for x := 0; x < len(user.Profiles); x++ {
			if user.Profiles[x].Type == r[i].EnumIndex() ||
				//Sysadmin can do everything
				user.Profiles[x].Type == enums.Sysadmin.EnumIndex() {
				checkRoles = true
				break
			}
		}
	}

	return checkRoles, err
}
