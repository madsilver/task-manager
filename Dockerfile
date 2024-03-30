FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./cmd ./cmd
COPY ./docs ./docs
COPY ./internal ./internal
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o task-manager cmd/api/main.go
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o task-manager cmd/api/main.go

FROM alpine:3.17
ARG DOCKER_USER=default_user
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/task-manager .
EXPOSE 8000
RUN addgroup -S "$DOCKER_USER" && adduser -S "$DOCKER_USER" -G "$DOCKER_USER"
USER $DOCKER_USER
CMD ["./task-manager"]