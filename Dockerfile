FROM golang:1.21

COPY . /go/src/app

WORKDIR /go/src/app/cmd/avito-tech

RUN go build -o avito main.go

EXPOSE 9090

CMD ["./avito"]