FROM golang:alpine3.23

WORKDIR /app

COPY . .

RUN go install github.com/air-verse/air@latest

RUN go mod tidy

EXPOSE 8080

CMD ["air", "--build.cmd", "go build -o /tmp/main /app/main.go", "--build.entrypoint", "/tmp/main", "--build.exclude_dir", "tmp", "--log.silent", "true"]