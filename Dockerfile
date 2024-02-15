FROM golang:1.21.3 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o microblog cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/microblog /app/microblog
COPY config/docker.yml /app/docker.yml
WORKDIR /app
CMD ["./microblog", "-config=docker.yml"]