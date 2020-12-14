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

docker-run-vm:
	@docker rm --force weebinar
	@docker rmi --force weebinar:latest
	@docker build --tag weebinar:latest .
	@docker run --restart unless-stopped --name weebinar -p 5000:5000 --detach --net host weebinar:latest