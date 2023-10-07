# build stage
FROM golang:1.21-alpine as builder

WORKDIR /app

# download modules as distinct layer
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# build code
COPY *.go ./
RUN go build -ldflags="-s -w" -o aka-redirector ./...

# runtime stage
FROM alpine

WORKDIR /app
COPY --from=builder /app/aka-redirector ./

EXPOSE 3000

CMD ["./aka-redirector"]