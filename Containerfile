FROM ghcr.io/hybridgroup/opencv:4.11.0 AS artifacts

ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin:/usr/local/go/bin

WORKDIR /app
COPY go.mod go.mod
RUN go mod download

COPY . .
RUN go build -v -x -ldflags "-w -s" -o main cmd/main.go
