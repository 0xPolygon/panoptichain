FROM golang:1.25-alpine AS builder
WORKDIR /app

RUN apk add --no-cache make

# Keep dependencies on their own layer to benefit from Docker's build cache
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the rest of the application code
COPY . .

ENV CGO_ENABLED=0 GOOS=linux
RUN make

FROM alpine:3.22
COPY --from=builder /app/out/panoptichain /panoptichain
COPY config.yml /etc/panoptichain/config.yml
RUN apk add --no-cache ca-certificates
CMD ["/panoptichain"]
