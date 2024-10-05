APP = $(error You must provide APP.)

build:
	CGO_ENABLED=0 go build -v -ldflags='-s -w' -trimpath -o bin/${APP} cmd/${APP}/main.go
	docker compose build app --build-arg APP_NAME=${APP}

run: build
	docker compose up

stop:
	docker compose down -v