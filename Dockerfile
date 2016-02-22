FROM golang:1.6

VOLUME ["/go/src/github.com/Nyarum/noterius"]

WORKDIR /go/src/github.com/Nyarum/noterius

ADD . /go/src/github.com/Nyarum/noterius

RUN go get -u github.com/cespare/reflex
RUN go get -u github.com/tools/godep
RUN godep restore

EXPOSE 1973

ENTRYPOINT ["reflex -s -- go run main.go"]