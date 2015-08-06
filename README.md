# Noterius - emulator for Pirate King Online

This emulator based on:
- PostgreSQL
- Docker
- fswatch

### Installation

- Download project
```
$ go get -u github.com/Nyarum/noterius
```

- Go to directory
```
$ cd $GOPATH/src/github.com/Nyarum/noterius
```

- Run build docker image
```
$ docker build -t noterius .
```

- Start project in docker
```
$ docker run -d -p 1973:1973 noterius fswatch
```