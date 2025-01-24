FROM golang:1.23.0 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o firefighting-operations-nland
RUN chmod +x /app/firefighting-operations-nland

FROM alpine:latest
WORKDIR /root
COPY --from=builder /app/firefighting-operations-nland /root/firefighting-operations-nland
CMD ["./firefighting-operations-nland"]