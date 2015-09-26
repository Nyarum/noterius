FROM golang:1.5.1

VOLUME ["/go/src/github.com/Nyarum/noterius"]

WORKDIR /go/src/github.com/Nyarum/noterius

ADD . /go/src/github.com/Nyarum/noterius

RUN go get -u gopkg.in/yaml.v2
RUN go get -u github.com/cespare/reflex
RUN go get -u github.com/lib/pq
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/Sirupsen/logrus

EXPOSE 1973

ENTRYPOINT ["fswatch"]