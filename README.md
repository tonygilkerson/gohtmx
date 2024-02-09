# gohtmx

Simple Go webserver serving up htmx

## Local testing

```sh
go run cmd/gohtmx/main.go
open http://localhost:8080/
```

## Deploy to test cluster

This can be use to test the app's chart

## Create Cluster

This requires port 8080 and 8443 to be available on the host.

```sh
podman machine init --cpus=4 --memory=4000 // adjust resources as needed
kind create cluster --config kind-ingress.yaml
```

## Install Ingress Controller

```sh
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s

# Patch the controller to listen on ports 8080 and 8443 in place of 80 and 443
kubectl -n ingress-nginx patch svc ingress-nginx-controller --type merge --patch-file nginx-patch.yaml
```

## Hit ap

```sh
open http://gohtmx.127.0.0.1.nip.io:8080/
```
