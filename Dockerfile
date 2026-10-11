FROM golang:1.20 AS builder
WORKDIR /app
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o my_company_app main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/my_company_app .
CMD ["./my_company_app"]
