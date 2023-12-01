# Build excutable
FROM golang:1.21 AS build-stage

WORKDIR /app
COPY . .

# Choose your platform
RUN GOARCH=arm  GOOS=linux go build -o /app/bin/wol-tg-bot

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS release-stage

WORKDIR /app

COPY --from=build-stage /app/bin .

ENTRYPOINT ["./wol-tg-bot"]