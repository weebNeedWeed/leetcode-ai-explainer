FROM golang:1.24.2-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# This command is commonly used in Docker containers based on Alpine Linux
# to ensure that applications within the container can make secure HTTPS connections to external services.
RUN apk add -U --no-cache ca-certificates

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./output ./cmd/api/

FROM scratch

WORKDIR /app

EXPOSE 9090

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/output ./

CMD ["./output"]