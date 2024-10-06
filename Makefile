APP = $(error You must provide APP.)

build:
	CGO_ENABLED=0 go build -v -ldflags='-s -w' -trimpath -o bin/${APP} cmd/${APP}/main.go

start-db:
	docker compose up -d db

run: build
	docker compose stop app || true
	docker compose rm app -vsf || true
	docker compose build app --build-arg APP_NAME=${APP}
	docker compose up app

stop:
	docker compose down -v --remove-orphans --rmi local