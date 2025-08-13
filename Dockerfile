FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY main.go .
RUN go mod init myimage && \
    go build -o myimage main.go

FROM alpine:3.22
RUN apk --no-cache add ca-certificates=20250619-r0
WORKDIR /root/
COPY --from=builder /app/myimage .

EXPOSE 80
CMD ["./myimage"]