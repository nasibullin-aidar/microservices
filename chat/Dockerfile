FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY ../go.mod ../go.sum ./

COPY ./chat/chat.go .

RUN GOOS=linux go build -v -o chat chat.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/chat .

EXPOSE 8080

CMD ["./chat"]