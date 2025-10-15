# validator-sim-infra/sim/Dockerfile
FROM golang:1.23-alpine

WORKDIR /app

# Copy your Go source code
COPY . .
RUN  go mod download 

# Build the exporter binary
RUN ls -la && go build -o validator-exporter validator.go

#check Health
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
    CMD wget --spider -q http://localhost:2112/metrics || exit 1
# Expose metrics port
EXPOSE 2112

# Run the exporter
CMD ["./validator-exporter"]

