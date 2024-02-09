# Development Notes

Notes for how to develop and test the about app on a local kind cluster.

## Setup

### Podman Machine

```sh
podman machine init --cpus=4 --memory=4000
```

### Create Cluster

This requires port 8080 and 8443 to be available on the host.

```sh
kind create cluster --config kind-ingress.yaml
```

### Install Ingress Controller

```sh
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s

# Patch the controller to listen on ports 8080 and 8443 in place of 80 and 443
kubectl -n ingress-nginx patch svc ingress-nginx-controller --type merge --patch-file nginx-patch.yaml
```

## Cluster

Once the cluster is ready with the ingress installed and configured you can use skaffold to deploy the about app.

To create an image pull secret first go [here](https://git.act3-ace.com/ace/about/-/settings/repository) and create a deploy token with `read_registry` access and save the name and

```sh
USER="<replace me>"
PASS="<replace me>"
EMAIL="<replace me>"
kubectl create secret docker-registry ace-ips --docker-server=reg.git.act3-ace.com --docker-username=$USER --docker-password=$PASS --docker-email=$EMAIL
```

Then use skaffold to deploy

```sh
skaffold dev --default-repo reg.git.act3-ace.com/ace/about 
# or
skaffold run --default-repo reg.git.act3-ace.com/ace/about 
```

Then open the following in your browser to hit the about endpoints.

```sh
open http://about.127.0.0.1.nip.io:8080/env/version/badge
```

## Workstation

To manually build the container

```sh
export KO_DOCKER_REPO="reg.git.act3-ace.com/ace/about"
export VERSION="v0.0.0-dev"

ko build ./cmd/about
```

To run the app on your workstation

```sh
export KO_DOCKER_REPO="reg.git.act3-ace.com/ace/about"
export VERSION="v0.0.0-dev"

export ABOUT_CONTEXT_FILE="$PWD/internal/env/test/ace-context.yaml"
export ABOUT_VERSION_BADGE_TEMPLATE_FILE="$PWD/internal/env/test/version-badge-template.gotmpl"
go run cmd/about/main.go
```
