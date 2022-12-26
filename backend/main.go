package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"github.com/kleberalves/problemCompanyApp/backend/user/controller"
	"github.com/kleberalves/problemCompanyApp/backend/user/repository"
	"github.com/kleberalves/problemCompanyApp/backend/user/service"
)

func main() {
	router := gin.Default()

	db := schema.Connect()
	schema.AutoMigrations()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	controller.NewUserController(router, userService)

	router.Run("localhost:8080")

}
