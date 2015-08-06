FROM golang:1.5beta3

VOLUME ["/go/src/github.com/Nyarum/noterius"]

WORKDIR /go/src/github.com/Nyarum/noterius

ADD . /go/src/github.com/Nyarum/noterius

RUN go get -u gopkg.in/yaml.v2
RUN go get -u github.com/codeskyblue/fswatch