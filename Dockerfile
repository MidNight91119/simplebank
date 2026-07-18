# Build stage
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run Stage
FROM alpine:3.22
WORKDIR /app 
COPY --from=builder /app/main .

EXPOSE 8080
CMD [ "/app/main" ]