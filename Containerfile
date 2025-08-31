FROM ghcr.io/hybridgroup/opencv:4.12.0 AS artifacts

ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin:/usr/local/go/bin

RUN apt update && apt install -y libtbb-dev

WORKDIR /app
COPY go.mod go.mod
RUN go mod download

COPY . .
RUN go build -v -x -ldflags "-w -s" -o main cmd/main.go
