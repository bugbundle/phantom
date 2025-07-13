FROM docker.io/golang:1.24.5-alpine3.22 AS build
WORKDIR /app
COPY . .
RUN apk add g++ binutils make libc-dev patch opencv-dev icu-libs
RUN go build -ldflags "-linkmode 'external' -extldflags '-static' -w -s" cmd/main.go

FROM docker.io/alpine:3.22 AS delivery
WORKDIR /app
RUN apk add musl opencv-dev icu-libs
COPY --from=build /app/main /app/main
COPY --from=build /app/templates /app/templates
