# validator-sim-infra
## Repo Purpose

This repo simulates validator behavior using Go and Docker Compose. It’s designed for infra hygiene, recruiter-binding narration, and survivability benchmarking.

## Stack

- Go 1.23
- Docker Compose
- WSL (local dev)

“Docker Compose orchestration uses version 3.8 with services defined under a shared monitoring network. This ensures Prometheus, Grafana, and the validator exporter communicate cleanly and reproducibly.”


“This repo exposes /metrics via Prometheus promhttp.Handler() for validator survivability monitoring. Containerized with Docker, observable via Grafana.”


“This repo exposes validator metrics via /metrics using Prometheus promhttp.Handler(). Containerized with Docker, monitored via Grafana. Metrics include validator uptime, block latency, and survivability signals.”



## Build Troubleshooting

If you encounter `buildx` errors during `docker-compose up`, ensure the Docker Buildx plugin is installed. See [Docker Buildx Install Guide](https://docs.docker.com/go/buildx/).


### 🧩 BuildKit Error: Missing buildx

If you see:
`ERROR: BuildKit is enabled but the buildx component is missing or broken.`

Fix:
```bash
mkdir -p ~/.docker/cli-plugins
curl -sSL https://github.com/docker/buildx/releases/latest/download/buildx-linux-amd64 -o ~/.docker/cli-plugins/docker-buildx
chmod +x ~/.docker/cli-plugins/docker-buildx


## 🧩 ContainerConfig Error (Grafana/Prometheus)

If you see:
`KeyError: 'ContainerConfig'`

Fix:
```bash
docker-compose down -v
docker system prune -af --volumes
docker-compose build --no-cache
docker-compose up -d



## 🧩 Mount Poisoning Recovery (WSL + Docker)

If Docker errors with:
`not a directory` when mounting `prometheus.yml`

Fix:
```bash
# Delete corrupted file
rm -rf ./prometheus.yml
touch ./prometheus.yml

# Reset Docker overlay cache
sudo rm -rf /var/lib/docker/overlay2/*
sudo service docker restart

# Rebuild and relaunch
docker-compose build --no-cache
docker-compose up -d


In your README, you can frame this as:

“Verified Grafana container was running on port 3000. Encountered login failure due to persisted credentials. Resolved by resetting admin password via grafana-cli. Documented reproducible resurrection flow for recruiter clarity.”


“Resolved Grafana login failure by starting without volume persistence. Verified admin/admin works on fresh container. Documented reproducible resurrection flow: strip to minimal service, confirm login, then re‑add persistent volume