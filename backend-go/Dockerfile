# --- build stage ---
FROM golang:1.24-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# go-fitz vendors static MuPDF libs for linux/amd64 and linux/arm64 (glibc),
# so no system MuPDF install is needed here — just a C compiler, which this
# base image already provides.
ENV CGO_ENABLED=1
RUN go build -o /out/server .

# needed for AWS SDK's HTTPS calls to Bedrock
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /out/server ./server
RUN mkdir -p uploads

EXPOSE 8080
 
ENTRYPOINT ["./server"]
