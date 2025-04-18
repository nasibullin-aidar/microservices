FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY ../go.mod ../go.sum ./

COPY ./auth/auth.go .

RUN GOOS=linux go build -v -o auth auth.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/auth .

EXPOSE 8080

CMD ["./auth"]