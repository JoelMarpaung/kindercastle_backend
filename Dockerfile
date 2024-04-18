FROM golang:1.22 as builder

WORKDIR /app
COPY . /app

RUN go install github.com/swaggo/swag/cmd/swag@latest && swag fmt && swag init -g internal/app/server/routes.go --parseDependency true --parseInternal true
RUN go mod download && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -installsuffix cgo -o kindercastle_backend main.go

FROM busybox
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/kindercastle_backend ./
CMD ["./kindercastle_backend", "serve-http"]
