FROM golang:alpine AS builder

ARG VERSION

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/locations-service

COPY . .

COPY go.mod go.sum ./
RUN go mod download

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -ldflags "-X main.Version=$VERSION" -o ./out/locations-service main.go

# Start fresh from a smaller image
FROM alpine:latest

RUN apk add ca-certificates

RUN apk add openssh-client

COPY --from=builder /tmp/locations-service/out/locations-service /app/locations-service

# This container exposes port 8091 to the outside world
EXPOSE 8081
EXPOSE 5051

CMD ["/app/locations-service","run-rest-server"]