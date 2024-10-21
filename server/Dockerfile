FROM golang:1.23.2 as builder

WORKDIR /app

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY ./src .

RUN go build -o expenses-server ./

FROM ubuntu:jammy

WORKDIR /app

COPY --from=builder /app/expenses-server .

EXPOSE 8080

CMD ["./expenses-server"]
