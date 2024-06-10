FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pismo ./cmd/pismo/main.go

EXPOSE 8080

CMD ["./pismo"]