package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_dataSource "github.com/kleberalves/problemCompanyApp/backend/datasources"
	httphandler "github.com/kleberalves/problemCompanyApp/backend/services/http-handler"

	"github.com/kleberalves/problemCompanyApp/backend/schema"

	_productController "github.com/kleberalves/problemCompanyApp/backend/product/controller"
	_productRepo "github.com/kleberalves/problemCompanyApp/backend/product/repository"
	_productService "github.com/kleberalves/problemCompanyApp/backend/product/service"

	_userController "github.com/kleberalves/problemCompanyApp/backend/user/controller"
	_useRepo "github.com/kleberalves/problemCompanyApp/backend/user/repository"
	_useService "github.com/kleberalves/problemCompanyApp/backend/user/service"

	_purchaseController "github.com/kleberalves/problemCompanyApp/backend/purchase/controller"
	_purchaseRepo "github.com/kleberalves/problemCompanyApp/backend/purchase/repository"
	_purchaseService "github.com/kleberalves/problemCompanyApp/backend/purchase/service"

	_profileController "github.com/kleberalves/problemCompanyApp/backend/profile/controller"
	_profileRepo "github.com/kleberalves/problemCompanyApp/backend/profile/repository"
	_profileService "github.com/kleberalves/problemCompanyApp/backend/profile/service"

	_credentialController "github.com/kleberalves/problemCompanyApp/backend/credential/controller"
	_credentialService "github.com/kleberalves/problemCompanyApp/backend/credential/service"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env not found.")
	}

	router := gin.Default()
	router.Use(httphandler.CorsMiddleware(false))

	db := _dataSource.NewPostGresDataSource().Connect()
	schema.Setup(db)

	userRepo := _useRepo.NewUserRepository(db)

	// Credential does not need self-repo because the password has been stored on the User.
	// Also needs to check for User existence by Email
	credentialService := _credentialService.NewCredentialService(userRepo)
	_credentialController.NewCredentialController(router, credentialService)

	userService := _useService.NewUserService(userRepo)
	_userController.NewUserController(router, userService, credentialService)

	productRepo := _productRepo.NewProductRepository(db)
	productService := _productService.NewProductService(productRepo)
	_productController.NewProductController(router, productService, credentialService)

	purchaseRepo := _purchaseRepo.NewPurchaseRepository(db)
	purchaseService := _purchaseService.NewPurchaseService(purchaseRepo)
	_purchaseController.NewPurchaseController(router, purchaseService, credentialService)

	profileRepo := _profileRepo.NewProfileRepository(db)
	profileService := _profileService.NewProfileService(profileRepo)
	_profileController.NewProfileController(router, profileService, credentialService)

	hostPort := os.Getenv("HOST_PORT")

	fmt.Println("Application port: " + hostPort)
	router.Run("localhost:" + hostPort)

}
