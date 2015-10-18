# Noterius - emulator for Pirate King Online
![](https://travis-ci.org/Nyarum/noterius.svg?branch=master)

Next-gen version with new architecture.

This emulator based on:
- PostgreSQL
- Docker
- fswatch
- rocker-compose
- LUA (Coming soon..)

### What does work?
- Authorization
- Exit
- Keep alive connect

### Installation

#### Docker - from latest sources

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
$ docker build -t nyarum/noterius .
```

- Start project in docker
```
$ rocker-compose run -attach
```

#### Docker - from stable releases

- Pull docker image
```
$ docker pull nyarum/noterius
```

- Start project in docker
```
$ rocker-compose run -attach
```

#### Use binary without docker

- Download binary from [releases page](https://github.com/Nyarum/noterius/releases)

- Install PostgreSQL

- Run binary and edit a config for your settings

### Trailer for game

[![Cool trailer for TOP ;)](http://img.youtube.com/vi/0l1TWRR5KuI/0.jpg)](http://www.youtube.com/watch?v=0l1TWRR5KuI)