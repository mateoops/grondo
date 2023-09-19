FROM golang:1.21-alpine

RUN apk add --no-cache gcc

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o bin .

EXPOSE 8080

ENTRYPOINT ["/app/bin"]