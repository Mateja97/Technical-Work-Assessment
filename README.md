# Alluvial Task - Ethereum Balance Proxy Service

This project is a Golang-based service that proxies the `eth_getBalance` RPC endpoint from the Ethereum execution layer. It is designed to be highly available (HA) and supports multiple Ethereum clients behind the proxy. The service also exposes Prometheus metrics and provides liveness and readiness HTTP endpoints for Kubernetes integration.

## Prerequisites

- Go 1.23.6
- Docker
- Kubernetes cluster (e.g., Minikube, GKE, EKS) - Minikube example shown here
- Ethereum client API keys (e.g., Infura, Alchemy, Tenderly)

## Configuration

The service is configured using environment variables. The following environment variables are required:

- `ETH_CLIENTS`: Comma-separated list of Ethereum client URLs.
- `SERVER_ADDRESS`: Address and port for the main service (e.g., `:8080`).

## Running steps:
- Set your docker-hub-username instead of <DOCKER_HUB_USERNAME> inside .kube/deployment.yaml
- Start kubernetes cluster:
```sh
minikube start
```
- Build image:
```sh
docker build -t  docker-hub-username/alluvial-task:latest .
```
- Push image:
```sh
docker push docker-hub-username/alluvial-task:latest 
```
- Apply kubernetes config:
```sh
 kubectl apply -f .kube/pod.yaml 
 kubectl apply -f .kube/service.yaml 
```

## Testing steps:
- Port forward:
```sh
 kubectl port-forward service/alluvial-task 8080:8080
```
- Execute getBalance request
```sh
curl -X GET localhost:8080/getBalance/<wallet-address>
```

## Metrics
```sh
curl -X GET localhost:8080/metrics
```