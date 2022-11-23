# syntax=docker/dockerfile:1

## Build
FROM golang:alpine3.16 as build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/*.go ./
COPY Makefile ./
COPY pkg/ ./pkg

# uncoment to generate new documentation
# RUN apk add --no-cache git
# RUN go install github.com/go-swagger/go-swagger/cmd/swagger@v0.29.0
# RUN swagger generate spec -o /swagger.json --scan-models
RUN go build -o /userapi

## Deploy
FROM alpine:3.16
WORKDIR /
# uncomment for fres docs
# COPY --from=build /swagger.json /docs/
COPY docs/swagger.json /docs/
COPY --from=build /userapi /userapi
EXPOSE 8200
ENTRYPOINT [ "/userapi" ]
