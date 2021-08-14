swagnrun: 
	swag init --parseDependency -g server/server.go && go run main.go

swag: 
	swag init --parseDependency -g server/server.go

run: 
	go run .
