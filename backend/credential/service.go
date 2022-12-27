package credential

import (
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

type Service interface {
	Login(user schema.UserCredential) (schema.Credential, error)
}
