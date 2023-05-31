# GRPC Load Balancing Kubernetes :tada:

## Setup

- Install kubernetes (docker desktop, minukube, ...)
- GRPC:

```shell
# Macos
brew install protobuf

# I’ll make sure it’s installed:
protoc --version

# Install the protocol compiler plugins for Go using the following commands:
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# Update your PATH so that the protoc compiler can find the plugins:
export PATH="$PATH:$(go env GOPATH)/bin"
```

```bash
# Apply confimap
kubectl apply -f ./k8s/configmap.yaml

# Apply server-grpc deployment
kubectl apply -f ./k8s/server-deployment.yaml

# Apply client-grpc deployment
kubectl apply -f ./k8s/client-deployment.yaml
```

## How to check the results

1. Go inside the pod to see

```bash
kubectl exec -it <pod name> -- bash
```

2. Install [k9s](https://k9scli.io/topics/install/)

## Demo is still in progress (`WIP`) :rocket:
