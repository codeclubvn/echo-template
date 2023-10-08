package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
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
	//bootstrap.Run()

	fmt.Println("Hello World")
	// Tạo một kênh để lắng nghe tín hiệu Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)

	// Sử dụng vòng lặp vô hạn để duy trì chương trình
	for {
		select {
		case <-c:
			// Nhận tín hiệu Ctrl+C, thoát ra khỏi vòng lặp
			fmt.Println("\nNhận tín hiệu Ctrl+C. Kết thúc chương trình.")
			return
		}
	}
}
