FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin .

EXPOSE 8080

ENTRYPOINT ["/app/bin"]