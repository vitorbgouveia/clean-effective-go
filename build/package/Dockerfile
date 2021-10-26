FROM golang

WORKDIR /go/src/cashback-sales-consumer

ADD . /go/src/cashback-sales-consumer

RUN go mod tidy
RUN gofmt -w ./**/*

EXPOSE 5000

CMD ["sh", "-c", "go run cmd/cashback-sales-service/main.go"]
