package schema

func AutoMigrations() {

	db := Connect()

	err := db.AutoMigrate(&User{}, &Product{}, &Purchase{}, &Profile{})
	if err != nil {
		return
	}
}
