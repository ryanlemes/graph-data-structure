ARG GO_VERSION=1.16.6

# first stage: Build
FROM golang:${GO_VERSION}-alpine as builder

# Set env variables to build on linux container
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod .

# Run to download the project dependencies
RUN go mod download

COPY . .

# Build the project existed in cmd/cli folder
RUN go build -o graph ./cmd/cli

# ---
# Final stage: using the scratch image to run the application
FROM scratch AS final

# Import the compiled executable from the first stage.
COPY --chown=0:0 --from=builder /app/graph /cli

ENTRYPOINT ["/cli"]