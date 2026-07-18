# Build stage
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run Stage
FROM alpine:3.22
WORKDIR /app 
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY db/migration ./migration
RUN chmod +x /app/start.sh

ARG TARGETARCH
RUN wget -qO- https://github.com/golang-migrate/migrate/releases/download/v4.19.1/migrate.linux-${TARGETARCH}.tar.gz | tar xz -C /usr/local/bin migrate

EXPOSE 8080
ENTRYPOINT [ "/app/start.sh" ]
CMD [ "/app/main" ]