# Build stage
FROM --platform=$BUILDPLATFORM golang:1.25 AS builder

WORKDIR /app

# Copy go mod files first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} \
    go build -trimpath -ldflags="-s -w" -o weight-controller ./cmd

# Runtime stage
FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/weight-controller /weight-controller

USER 65532:65532

ENTRYPOINT ["/weight-controller"]
