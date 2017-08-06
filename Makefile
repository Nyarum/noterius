
NAME=noterius

setup:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/mattes/migrate

generate:
	go generate ./...

build: generate
	go build -o ${NAME} main.go

run: build
	./${NAME}

migrate_up:
	migrate -path ./migrations -database 'postgres://noterius:noterius@localhost:5455/noterius?sslmode=disable' up
	
migrate_down:
	migrate -path ./migrations -database 'postgres://noterius:noterius@localhost:5455/noterius?sslmode=disable' down