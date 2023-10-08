FROM golang:1.21-alpine as build-env

# cache dependencies first
WORKDIR /app
COPY go.mod /app
COPY go.sum /app
RUN go mod download

COPY . /app

# build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/cmd/main ./cmd/main.go

FROM alpine:latest

WORKDIR /app
COPY ./config/config.yml /app/config/config.yml

COPY --from=build-env /app/cmd/main /app/cmd/main

ENTRYPOINT ["/app/main"]


