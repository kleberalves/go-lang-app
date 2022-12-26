package schema

import "gorm.io/gorm"

func AutoMigrations(db *gorm.DB) {

	db.AutoMigrate(&User{}, &Product{}, &Purchase{}, &Profile{})

	if db.Migrator().HasConstraint(&Profile{}, "idx_profiles_user_id") {
		db.Migrator().DropConstraint(&Profile{}, "idx_profiles_user_id")
	}

}
