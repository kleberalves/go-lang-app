package schema

import (
	"fmt"
	"os"

	"github.com/kleberalves/problemCompanyApp/backend/enums"
	"github.com/kleberalves/problemCompanyApp/backend/services/security"

	"gorm.io/gorm"
)

func Setup(db *gorm.DB) {

	db.AutoMigrate(&User{}, &Product{}, &Purchase{}, &Profile{})

	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPwd := os.Getenv("ADMIN_PWD")

	if adminEmail != "" && adminPwd != "" {

		var user User
		err := db.Model(&User{}).
			Where("email = ?", adminEmail).
			First(&user).Error

		if user.Email == "" || (err != nil && err.Error() == "record not found") {

			adminPwd, err = security.HashPassword(adminPwd)
			if err == nil {

				user = User{
					FirstName: "Admin",
					LastName:  "Sys",
					Email:     adminEmail,
					Password:  adminPwd,
					Profiles: []Profile{{
						Type: enums.Sysadmin.EnumIndex(),
					}},
				}

				err = db.Model(&User{}).
					Create(&user).
					Error

				if err != nil {
					panic("can't create user admin: " + err.Error())
				} else {
					fmt.Println("User Sysadmin ", adminEmail, " has been created.")
				}
			}
		}
	}

}
