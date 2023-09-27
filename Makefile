init_and_run: swagger_init run_source_code

swagger_init:
	swag init

run_source_code:
	go run main.go
