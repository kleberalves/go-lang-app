package user

import (
	"github.com/kleberalves/problemCompanyApp/backend/enums"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

// Repository represent the article's repository contract
type Repository interface {
	FindAll() (res []schema.UserRead, err error)
	Create(user schema.User) (schema.User, error)
	AssociateProfile(userId int, typ enums.TypeUser) schema.Profile
	RemoveProfile(userId int, typ enums.TypeUser)
	// GetByID(ctx context.Context, id int64) (*models.Article, error)
	// GetByTitle(ctx context.Context, title string) (*models.Article, error)
	// Update(ctx context.Context, ar *models.Article) error
	// Store(ctx context.Context, a *models.Article) error
	// Delete(ctx context.Context, id int64) error
}
