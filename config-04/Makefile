## Docker image name
IMG := eduardobcolombo/
VERSION := 1.0 # $(shell git rev-parse --short HEAD)
CLUSTER := colombostips-kube
NAMESPACE := colombostips
GO_VERSION := 1.20
LOCAL_PATH=/app

##################
### Build section
##################
## Build a docker image for the server service.
build-img:
	docker build \
	-f ./docker/Dockerfile.server \
	-t $(IMG)server:$(VERSION) \
	--build-arg BUILD_REF=$(VERSION) \
	--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
	.
.PHONY: build-img

## Build a binary file for the server service.
build-bin: 
	@docker run --rm \
		-v ${PWD}:$(LOCAL_PATH) \
		-w $(LOCAL_PATH) -e GOOS=linux -e GOARCH=amd64 \
		golang:$(GO_VERSION) \
		go build -o bin/server.bin ./main.go
.PHONY: build-bin

##################
### KinD - Kubernetes in Docker section.
##################
## Shortcut to init KinD with the images.
kind-init: kind-create kind-config build-img kind-load kind-apply kind-apply-secrets 
.PHONY: kind-init

kind-reload: build-img kind-load kind-apply kind-apply-secrets kind-restart
.PHONY: kind-reload

kind-restart:
	kubectl rollout restart deployment server --namespace=$(NAMESPACE)
.PHONY: kind-restart

## Creates a KinD cluster locally.
kind-create: 
	@kind create cluster --name $(CLUSTER)
.PHONY: kind-create

## Apply secrets based on the .env file
kind-apply-secrets:
	@kubectl create secret generic server-secrets \
		--from-env-file=./deploy/server.env \
		--namespace $(NAMESPACE) --cluster kind-$(CLUSTER) --dry-run=client -o yaml \
	> ./k8s/server-secrets.yaml
	@kubectl apply -f ./k8s/server-secrets.yaml --namespace $(NAMESPACE) --cluster kind-$(CLUSTER)
.PHONY: kind-apply-secrets

## Set the current cluster to the context and load env variables.
kind-config:
	@kubectl config set-context --current --namespace $(NAMESPACE) --cluster kind-$(CLUSTER)
.PHONY: kind-config

## Load docker built images to the KinD cluster.
kind-load:
	kind load docker-image $(IMG)server:$(VERSION) --name $(CLUSTER)
.PHONY: kind-load

## Apply the k8s folder with yaml files to the cluster.
kind-apply: 
	@kubectl apply -f ./k8s/namespace --cluster kind-$(CLUSTER)
	@kubectl apply -f ./k8s/deployment --cluster kind-$(CLUSTER)
.PHONY: kind-apply

# Delete all resources hostedn in the ./k8s
kind-delete:
	@kubectl delete -f ./k8s/namespace --cluster kind-$(CLUSTER)
.PHONY: kind-delete

# Delete the cluster.
kind-clean:
	@kind delete cluster --name $(CLUSTER)
.PHONY: kind-clean

# Monitoring with watch
kind-watch:
	@watch kubectl get all --namespace $(NAMESPACE)
.PHONY: kind-watch

expose:
	@kubectl port-forward service/server 8080
	&&
.PHONY: expose

