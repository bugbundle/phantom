FROM ghcr.io/hybridgroup/opencv:4.9.0 AS artifacts

RUN wget https://go.dev/dl/go1.25.0.linux-arm64.tar.gz && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.25.0.linux-arm64.tar.gz

ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin:/usr/local/go/bin

WORKDIR /app
COPY go.mod go.mod
RUN go mod download

COPY . .
RUN go build -v -x -ldflags "-w -s" -o main cmd/main.go
