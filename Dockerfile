# Start from the official Golang image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Install necessary dependencies
RUN apk add --no-cache git

# Clone the repository
RUN git clone https://github.com/PrivixAI-labs/Privix-node
RUN 

# Download the Go modules
RUN go mod tidy && go mod download

# Build the Go binary
RUN go build 

# Expose necessary ports
EXPOSE 10001 8545

# Set up a volume for persistent data storage
VOLUME [ "/app/data" ]

# Define the command to run your blockchain node
CMD ["./neth", "server", "--data-dir=data", "--chain", "genesis.json", "--libp2p", "0.0.0.0:10001", "--nat", "0.0.0.0", "--jsonrpc", "0.0.0.0:8545", "--seal", "--block-gas-target", "5000000000"]
