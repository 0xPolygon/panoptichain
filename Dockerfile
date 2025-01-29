FROM golang:1.21-bullseye AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0 GOOS=linux
RUN make

FROM ubuntu:24.04
COPY --from=builder /app/out/panoptichain /panoptichain
COPY config.yml /etc/panoptichain/config.yml
RUN apt-get update && apt-get install -y ca-certificates && apt-get clean
EXPOSE 9090
CMD ["/panoptichain"]
