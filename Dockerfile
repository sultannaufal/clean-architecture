# syntax=docker/dockerfile:1

FROM golang:1.19-alpine as base
WORKDIR /builder

ENV GO111MODULE=on CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

# COPY *.go ./
COPY . .

RUN go build -o /builder/main /builder/main.go

# runner image
FROM gcr.io/distroless/static:latest
WORKDIR /app
COPY --from=base /builder/main main
COPY --from=base /builder/.env .env

EXPOSE 8000
CMD ["/app/main"]