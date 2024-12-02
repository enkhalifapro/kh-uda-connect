FROM golang:alpine AS builder

ARG VERSION

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/persons-service

COPY . .

COPY go.mod go.sum ./
RUN go mod download

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -ldflags "-X main.Version=$VERSION" -o ./out/persons-service main.go

# Start fresh from a smaller image
FROM alpine:latest

RUN apk add ca-certificates

RUN apk add openssh-client

COPY --from=builder /tmp/persons-service/out/persons-service /app/persons-service

# This container exposes port 8091 to the outside world
EXPOSE 8092

CMD ["/app/persons-service"]