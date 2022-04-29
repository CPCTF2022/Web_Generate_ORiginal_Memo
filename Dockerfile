# syntax = docker/dockerfile:1.3.0

FROM node:18.0-alpine3.14 AS build-client

WORKDIR /app

COPY client/package*.json ./
RUN npm ci
COPY client/ .
RUN NODE_ENV=production npm run build

FROM golang:1.17.9-alpine AS build-server

RUN apk add --update --no-cache git

WORKDIR /go/src/github.com/CPCTF2022/Web_Generate_ORiginal_Memo/server

COPY server/go.mod server/go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
  go mod download

COPY server/ ./

RUN go build -o gorm -ldflags "-s -w"

FROM alpine:3.15.4

WORKDIR /app

RUN apk --update --no-cache add tzdata \
  && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && apk del tzdata \
  && mkdir -p /usr/share/zoneinfo/Asia \
  && ln -s /etc/localtime /usr/share/zoneinfo/Asia/Tokyo

COPY --from=build-client /app/build/ ./dist/
COPY --from=build-server /go/src/github.com/CPCTF2022/Web_Generate_ORiginal_Memo/server/gorm ./gorm

ENTRYPOINT ./gorm
