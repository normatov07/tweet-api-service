FROM golang:1.22-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN apk --no-cache add gcc g++ make git

RUN mkdir -p /go/src/app

WORKDIR /go/src/app

COPY ./src/go.mod /go/src/app
COPY ./src/go.sum /go/src/app

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY ./src /go/src/app

RUN GOOS=linux CGO_ENABLED=1 GOARCH=amd64 go build -o main ./cmd/main.go

ENV TZ=Asia/Tashkent

CMD ["./main"]