# gohtmx

Simple Go webserver serving up htmx

## Local testing

```sh
export WEB_ROOT=$PWD
go run cmd/gohtmx/main.go
open http://localhost:8080/
```

## Testing with containers

```sh
# Create VM
KIND_EXPERIMENTAL_PROVIDER=podman
podman machine init --cpus=4 --memory=4000 // adjust resources as needed

# Build container
podman build -t gohtmx:dev .

# Run
podman run -it --rm -p 8080:8080 localhost/gohtmx:dev 
open http://localhost:8080/
```

## Deploy to test cluster

This can be use to test the app's chart

## Create Cluster

This requires port 8080 and 8443 to be available on the host.

```sh
# Create VM and cluster
KIND_EXPERIMENTAL_PROVIDER=podman
podman machine init --cpus=4 --memory=4000 // adjust resources as needed
kind create cluster --config kind-confg.yaml

# Create and load image onto cluster nodes so we don't need a registry
podman save -o .temp/gohtmx.tar localhost/gohtmx:dev
kind load image-archive .temp/gohtmx.tar 

# Verify
# If you want to see the image you just loaded on the kind node
$ podman machine ssh
core@localhost:~$  podman exec -it <replace-me-with-container-id> /bin/bash
root@kind-worker:/# crictl images  
```

## Deploy App

```sh
helm upgrade -i gohtmx charts/gohtmx
open http://gohtmx.127.0.0.1.nip.io:8080/
```
