# syntax=docker/dockerfile:1

FROM golang:1.19-alpine as base
WORKDIR /builder

ENV GO111MODULE=on CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /builder/main

# runner image
FROM gcr.io/distroless/static:latest
WORKDIR /app
COPY --from=base /builder/main /app/main
# COPY --from=base /builder/.env /app/.env

EXPOSE 8000
ENTRYPOINT ["/app/main"]