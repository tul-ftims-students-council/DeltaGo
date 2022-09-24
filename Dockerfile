FROM golang:1.18
WORKDIR /go/src/delta-go
COPY . .
RUN go build -o bin/server cmd/main.go
CMD ["./bin/server"]