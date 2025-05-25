FROM docker.io/golang:1.24.2-alpine3.21 AS build
WORKDIR /app
COPY . .
RUN apk add build-base musl-dev opencv-dev icu-libs --repository=https://dl-cdn.alpinelinux.org/alpine/edge/community
RUN go build cmd/main.go

FROM docker.io/alpine:3.21 AS delivery
WORKDIR /app
RUN apk add musl opencv-dev icu-libs --repository=https://dl-cdn.alpinelinux.org/alpine/edge/community
COPY --from=build /app/main /app/main
COPY --from=build /app/api/templates /app/api/templates
