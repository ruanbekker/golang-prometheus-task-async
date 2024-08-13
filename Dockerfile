# Builder stage
FROM golang:1.20-alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main .

# Start a new stage from alpine
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /

COPY --from=builder /main .

EXPOSE 8080

CMD ["/main"]
