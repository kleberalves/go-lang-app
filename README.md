# Problem Company App Test

## Backend setup

### Environment variables

#### Database Connection

	POSTGRES_USER=postgres
	POSTGRES_PASSWORD=postgres
	POSTGRES_DB=postgres
	POSTGRES_HOSTNAME=localhost
	POSTGRES_PORT=5432

#### Host port

    HOST_PORT=8080
    
#### Jwt secret
To enable JWT token

	JWT_SECRET=kda8aA5s2a93jsaZadJSDNals018jd
    
## Creating Sysadmin user

All access points are restricted by the user type (called Profiles), so to init the application, had to create a new user type called Sysadmin that can do everything. Just set these environment variables before starting the app:

    ADMIN_EMAIL=admin@problemcompany.org
    ADMIN_PWD=test123

### Postman collection

All endpoints are available in [this](https://api.postman.com/collections/4546858-3803d996-9c08-4125-8eba-9fe09f899a18?access_key=PMAT-01GN803M19S4ZCQ7TMBZHFRKDD) postman collection.

## Credential future

### Increase security with recaptcha on anonymous calls
    router.POST("/credential/login/:recaptcha", ctrl.Login)
    router.POST("/credential/login/activate/:recaptcha", ctrl.XXX)
    router.POST("/credential/reset-password-send/:email/:recaptcha", ctrl.XXX)
    router.POST("/credential/reset-password/:recaptcha", ctrl.XXX)

### Enable - On Time Password
    router.GET("/credential/generate-secret-otp", ctrl.XXX)
    router.POST("/credential/validate-otp", ctrl.XXX)
    router.POST("/credential/remove-otp", ctrl.XXX)

	
