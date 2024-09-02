FROM golang:1.21.5-alpine as build

WORKDIR /

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /geminicare ./cmd/app/main.go

FROM gcr.io/distroless/base-debian11 AS build-release

WORKDIR /

COPY --from=build /geminicare /geminicare

EXPOSE 1323

USER nonroot:nonroot

ENTRYPOINT ["/geminicare"]

