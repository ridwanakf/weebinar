# go build command
build:
	@echo " >> building binaries"
	@go build -v -o bin/weebinar cmd/weebinar/*.go

# go run command
run: build
	./bin/weebinar

# run all go:generate commands (eg. Mock files generator)
generate:
	@go generate ./...

# run with docker
docker-run:
	@docker-compose up -d

docker-stop:
	@docker-compose down