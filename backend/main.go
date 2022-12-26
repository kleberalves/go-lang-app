package main

import (
	"github.com/gin-gonic/gin"
	_dataSource "github.com/kleberalves/problemCompanyApp/backend/datasources"

	"github.com/kleberalves/problemCompanyApp/backend/schema"

	_productController "github.com/kleberalves/problemCompanyApp/backend/product/controller"
	_productRepo "github.com/kleberalves/problemCompanyApp/backend/product/repository"
	_productService "github.com/kleberalves/problemCompanyApp/backend/product/service"

	_userController "github.com/kleberalves/problemCompanyApp/backend/user/controller"
	_useRepo "github.com/kleberalves/problemCompanyApp/backend/user/repository"
	_useService "github.com/kleberalves/problemCompanyApp/backend/user/service"
)

func main() {
	router := gin.Default()

	db := _dataSource.NewPostGresDataSource().Connect()
	schema.AutoMigrations(db)

	userRepo := _useRepo.NewUserRepository(db)
	userService := _useService.NewUserService(userRepo)
	_userController.NewUserController(router, userService)

	productRepo := _productRepo.NewProductRepository(db)
	productService := _productService.NewProductService(productRepo)
	_productController.NewProductController(router, productService)

	router.Run("localhost:8080")

}
