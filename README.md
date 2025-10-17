# validator-sim-infra

A reproducible validator simulation stack with Prometheus metrics and Grafana visualization. Built for infra hygiene, audit clarity, and recruiter-ready signaling.

---

## 🧩 Repo Purpose

This repo simulates validator behavior using Go and Docker Compose. It’s designed for infra hygiene, recruiter-binding narration, and survivability benchmarking.

---

## 🐳 Stack

- Go 1.21
- Docker Compose
- WSL (local dev)

Docker Compose orchestration uses version 3.8 with services defined under a shared monitoring network. This ensures Prometheus, Grafana, and the validator exporter communicate cleanly and reproducibly.

---

## ⚙️ Dockerfile Overview

```Dockerfile
FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o validator-exporter validator.go
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
  CMD wget --spider -q http://localhost:2112/metrics || exit 1
EXPOSE 2112
CMD ["./validator-exporter"]

## 🔍 Prometheus Configuration

Prometheus scrapes metrics from the custom validator-exporter container every 15 seconds.

### `prometheus.yml`
```yaml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'validator-exporter'
    static_configs:
      - targets: ['validator-exporter:2112']

📊 Grafana Dashboard
![Dashboard Name: validator monitoring] (assets/validator.png)

Panels:

validator_reward_total → Time Series

validator_penalty_total → Bar Gauge

validator_slashed → Stat (0 = Not Slashed, 1 = Slashed)

validator_removal_total → Time Series