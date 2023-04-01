FROM golang:1.20-alpine as builder

VOLUME [ "/go" ]

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -o bakery-app \
  ./cmd/bakery-app/bakery-app.go

FROM --platform=linux/amd64 alpine

WORKDIR /app

COPY --from=builder /app/bakery-app /app

CMD [ "./bakery-app" ]
