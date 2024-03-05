# Build the application from source
FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o sky_high main.go

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS running-stage

WORKDIR /

ENV VERSION 1.0.0

COPY --from=build-stage /app/sky_high /sky_high

USER nonroot:nonroot

ENTRYPOINT ["/sky_high"]
