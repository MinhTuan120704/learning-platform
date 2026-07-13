# BUILD STAGE

FROM golang:1.26 AS BUILDER

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o identity ./cmd/api

#RUNTIME

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

RUN groupadd -r appuser && useradd -r -g appuser appuser

WORKDIR /app

COPY --from=BUILDER app/identity .

RUN chown appuser:appuser /app/identity

USER appuser

EXPOSE 8081

CMD [ "./identity" ]