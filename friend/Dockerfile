FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY ../go.mod ../go.sum ./

COPY ./friend/friend.go .

RUN GOOS=linux go build -v -o friend friend.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/friend .

EXPOSE 8080

CMD ["./friend"]