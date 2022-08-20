KIND_CLUSTER_NAME := grpc-num-1
IMAGES_VERSION := 1.0.0
OS   := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)

KUBERNETES_VERSION         := 1.23.0
KIND_VERSION               := 0.11.1

BIN_DIR := $(shell pwd)/bin

KUBECTL                 := $(abspath $(BIN_DIR)/kubectl)
KIND                    := $(abspath $(BIN_DIR)/kind)

KUBECTL_CMD := KUBECONFIG=./.kubeconfig $(KUBECTL)
KIND_CMD    := $(KIND) --name $(KIND_CLUSTER_NAME)

kubectl: $(KUBECTL)
$(KUBECTL):
	curl -Lso $(KUBECTL) https://storage.googleapis.com/kubernetes-release/release/v$(KUBERNETES_VERSION)/bin/$(OS)/$(ARCH)/kubectl
	chmod +x $(KUBECTL)

kind: $(KIND)
$(KIND):
	curl -Lso $(KIND) https://github.com/kubernetes-sigs/kind/releases/download/v$(KIND_VERSION)/kind-$(OS)-$(ARCH)
	chmod +x $(KIND)


deploy:	$(KUBECTL) $(KIND) 
	make kind_make_cluster
	make prep_ingress
	make kind_load
	make deploy_services

prep_ingress:
	./sh/prepare_ingress.sh
	$(KUBECTL_CMD) apply -f https://raw.githubusercontent.com/metallb/metallb/v0.10.2/manifests/namespace.yaml
	$(KUBECTL_CMD) apply -f https://raw.githubusercontent.com/metallb/metallb/v0.10.2/manifests/metallb.yaml
	$(KUBECTL_CMD) apply -f ./metallb.yaml
deploy_services:
	$(KUBECTL_CMD) apply -f ./services.yaml
	$(KUBECTL_CMD) get all -A
kind_make_cluster:
	$(KIND_CMD) delete cluster --name $(KIND_CLUSTER_NAME)  --kubeconfig ./.kubeconfig
	$(KIND_CMD) create cluster --image kindest/node:v${KUBERNETES_VERSION} --config ./kind.yaml  --kubeconfig ./.kubeconfig
kind_load:
	$(KIND_CMD) load --name $(KIND_CLUSTER_NAME) docker-image \
	grpc-microservice-example-sum:$(IMAGES_VERSION) \
	grpc-microservice-example-sub:$(IMAGES_VERSION) \
	grpc-microservice-example-mul:$(IMAGES_VERSION) \
	grpc-microservice-example-div:$(IMAGES_VERSION) \
	grpc-microservice-example-auth:$(IMAGES_VERSION) \
	grpc-microservice-example-api:$(IMAGES_VERSION) || true
list:
	$(KUBECTL_CMD) get all -A

clean:
	$(KIND_CMD) delete cluster --name $(KIND_CLUSTER_NAME)