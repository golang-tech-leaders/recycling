FROM golang:1.15-buster as builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download -x

COPY . ./

ENV CGO_ENABLED=0
RUN go build -v -o /bin/server cmd/recyclingd/main.go

FROM debian:buster-slim
RUN set -x && apt-get update && \
  DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates && \
  rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY config.yml ./
COPY --from=builder /bin/server ./

ENV PORT=8080
ENV DBPORT=15432

EXPOSE $PORT

CMD ["./server", "-config", "config.yml"]


