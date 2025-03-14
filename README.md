# Alluvial Task - Ethereum Balance Proxy Service

This project is a Golang-based service that proxies the `eth_getBalance` RPC endpoint from the Ethereum execution layer. It is designed to be highly available (HA) and supports multiple Ethereum clients behind the proxy. The service also exposes Prometheus metrics and provides liveness and readiness HTTP endpoints for Kubernetes integration.

## Features

- **High Availability (HA)**: The service is designed to handle multiple Ethereum clients, ensuring redundancy and fault tolerance.
- **Multiple Ethereum Clients**: Supports multiple Ethereum execution node clients (e.g., Infura, Alchemy, Tenderly).
- **Inconsistent Data Handling**: Implements a strategy to handle inconsistent data returned by different Ethereum clients.
- **Prometheus Metrics**: Exposes metrics at the `/metrics` endpoint for monitoring.
- **Liveness and Readiness Probes**: Provides HTTP endpoints for Kubernetes liveness and readiness checks.

## Prerequisites

- Go 1.19 or higher
- Docker
- Kubernetes cluster (e.g., Minikube, GKE, EKS)
- Ethereum client API keys (e.g., Infura, Alchemy, Tenderly)

## Configuration

The service is configured using environment variables. The following environment variables are required:

- `ETH_CLIENTS`: Comma-separated list of Ethereum client URLs.
- `SERVER_ADDRESS`: Address and port for the main service (e.g., `:8080`).
- `HEALTH_CHECK_ADDRESS`: Address and port for the health check endpoint (e.g., `:8081`).

Example configuration:



