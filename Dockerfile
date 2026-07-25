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
COPY db/migration ./db/migration
RUN chmod +x /app/start.sh

EXPOSE 8080
ENTRYPOINT [ "/app/start.sh" ]
CMD [ "/app/main" ]