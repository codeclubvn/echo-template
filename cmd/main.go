package main

import (
	"echo_template/bootstrap"
)

//	@title			Swagger Echo-Template Backend API
//	@version		2.0
//	@description	This is Echo-Template Backend API docs
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey	Authorization
// @in							header
// @name						Authorization
// @host						localhost:8010
// @BasePath					/v1/api
func main() {
	bootstrap.Run()
}
