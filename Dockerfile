FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o bin -a -ldflags '-linkmode external -extldflags "-static"' .

EXPOSE 8080

ENTRYPOINT ["/app/bin"]