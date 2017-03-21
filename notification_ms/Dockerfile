FROM golang

MAINTAINER Neill Giraldo <neillro24124h@gmail.com>


COPY . /go/src/notification_ms
WORKDIR /go/src/notification_ms

RUN go get github.com/ant0ine/go-json-rest/rest
RUN go get gopkg.in/gomail.v2

RUN go build

CMD ["notification_ms"]