FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o mcd-clone main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/mcd-clone/ .
COPY --from=builder /app/assets/data ./data

CMD ["/app/mcd-clone"]