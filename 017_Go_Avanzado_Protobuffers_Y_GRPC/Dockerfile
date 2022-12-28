FROM golang:alpine AS builder

RUN apk add --no-cache git
RUN mkdir /go/bin/app
RUN chmod +x /go/bin/app
RUN apk update && apk add openssh

WORKDIR /build

RUN apk --no-cache add ca-certificates && update-ca-certificates

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main ./cmd

WORKDIR /dist

RUN cp /build/main .
##RUN cp /build/etc/ssl/certs/ca-certificates.crt /ect/ssl/certs/

EXPOSE 5050

CMD ["/dist/main"]