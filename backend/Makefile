.PHONY: test cover lint build docker-build docker-run migrate

test:
	go test -v ./... -cover

cover:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

lint:
	golangci-lint run ./...

build:
	CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

docker-build:
	docker build -t jp-song-trainer-backend .

docker-run:
	docker run -p 8080:8080 \
		-e DATABASE_URL="postgresql://user:password@localhost:5432/jp_song_trainer?sslmode=disable" \
		-e JWT_SECRET="local-secret" \
		jp-song-trainer-backend

migrate:
	go run ./cmd/server migrate