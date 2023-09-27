init_and_run: swagger_init run_source_code

swagger_init:
	swag fmt -d ./presenter/controller
	swag init -d ./cmd,./ -o ./presenter/docs

run_source_code:
	go mod tidy
	go run ./cmd/main.go
