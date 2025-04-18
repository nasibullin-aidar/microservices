FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY ../go.mod ../go.sum ./

COPY ./user/user.go .

RUN GOOS=linux go build -v -o user user.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/user .

EXPOSE 8080

CMD ["./user"]