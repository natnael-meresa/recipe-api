FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app/bin/recipe cmd/main.go


FROM alpine

WORKDIR /app

COPY --from=builder /app/bin/recipe .

EXPOSE 9000

CMD ["./recipe"]