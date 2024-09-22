# Build the application from source
FROM golang:1.22.5 AS build-stage

EXPOSE 8080

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy Modules
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /printrequest

# Minimize Build
FROM gruebel/upx:latest as minimize-stage
COPY --from=build-stage /printrequest /printrequest.old

# Compress the binary and copy it to final image
RUN upx --best --lzma -o /printrequest /printrequest.old



# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian12 AS build-release-stage
# FROM debian:12 AS build-release-stage

WORKDIR /app

# copy build
COPY --from=minimize-stage /printrequest /app/printrequest

ENTRYPOINT ["/app/printrequest"]