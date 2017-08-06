
NAME=noterius

setup:
	go get -u github.com/kardianos/govendor
	go get -u gopkg.in/src-d/go-kallax.v1/...

	# Bug with init() function in the driver
	rm -rf ${GOPATH}/src/github.com/lib/pq

generate:
	rm -rf ./models/kallax.go
	go generate ./models

build:
	go build -o ${NAME} main.go

run: build
	./${NAME}

migrate_up:
	kallax migrate up --all --dsn 'noterius:noterius@localhost:5455/noterius?sslmode=disable'
	
migrate_down:
	kallax migrate down --steps 1 --dsn 'noterius:noterius@localhost:5455/noterius?sslmode=disable'

migrate_generate: generate
	kallax migrate --input ./models/ --out ./migrations --name initial_schema

	$(MAKE) migrate_up