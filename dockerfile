FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod init fullcycle

RUN go build -o main .

FROM scratch

WORKDIR /app

COPY --from=builder /app /app

CMD ["./main"]