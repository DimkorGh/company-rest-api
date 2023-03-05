# -------- APP COMMANDS ------------

docker.app.start:
	docker-compose build
	docker-compose up

# -------- FORMAT COMMANDS ------------

docker.format:
	make dockerBuild
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			api-build go fmt ./...

# -------- LINTER COMMANDS ------------

docker.linter.run:
	make dockerBuild
		@docker run \
				--rm \
				--volume "$(PWD)":/app \
				--workdir /app \
				api-build golangci-lint run

# -------- TESTS COMMANDS ------------

docker.test.unit:
	make dockerBuild
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			api-build go test -tags musl -short -count=1 ./...

docker.test.all:
	make dockerBuild
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			api-build go test -tags musl -count=1 ./...

docker.test.all.coverage.withView:
	make dockerBuild
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			api-build go test -tags musl -count=1 -v -coverprofile=profile.cov ./... ; go tool cover -func profile.cov && go tool cover -html=profile.cov

# -------- MOCK COMMANDS -------

docker.mock.generate:
	make dockerBuild
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			api-build mockgen -source="$(FILE)"

docker.swagger.generate:
	make dockerBuild
	@docker run \
			--rm \
			--volume "$(PWD)":/app \
			--workdir /app \
			api-build swag init -g /cmd/main/main.go

# -------- DOCKER BUILDS -------

dockerBuild:
	@docker build \
			--tag api-build \
			-f Dockerfile.test .
