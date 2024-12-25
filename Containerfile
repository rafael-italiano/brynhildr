### Step 1: Build stage
FROM docker.io/library/golang:alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o brynhildr ./cmd/brynhildr/

FROM scratch
COPY --from=builder /app/brynhildr /
ENTRYPOINT ["/brynhildr"]
