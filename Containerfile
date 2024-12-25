### Step 1: Build stage
FROM docker.io/library/golang:alpine as builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code (Go source files)
COPY . .
RUN ls
RUN pwd
# Build the Go application with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o brynhildr ./cmd/brynhildr/

### Step 2: Final stage
FROM scratch

# Copy only the binary from the build stage to the final image
COPY --from=builder /app/brynhildr /

# Set the entry point for the container
ENTRYPOINT ["/brynhildr"]
