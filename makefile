# ==============================================================================
# Building containers

all: api

api:
	docker build \
		-f ./Dockerfile \
		-t comments-api-amd64:1.0 \
		--build-arg VCS_REF=`git rev-parse HEAD` \
		--build-arg BUILD_DATE=`date -u +”%Y-%m-%dT%H:%M:%SZ”` \
		.

# ==============================================================================
# Running from within docker compose

run: up

up:
	docker-compose -f docker-compose.yml up --detach --remove-orphans

down:
	docker-compose -f docker-compose.yml down --remove-orphans

logs:
	docker-compose -f docker-compose.yml logs -f



# ==============================================================================
# Running from within k8s/dev

kind-up:
	kind create cluster --image kindest/node:v1.20.2 --name comments-api-cluster --config config/dev/kind-config.yml

kind-down:
	kind delete cluster --name comments-api-cluster

kind-load:
	kind load docker-image comments-api-amd64:1.0 --name comments-api-cluster

kind-services:
	kustomize build config/dev | kubectl apply -f -

kind-update: api
	kind load docker-image comments-api-amd64:1.0 --name comments-api-cluster
	kubectl delete pods -lapp=comments-api

kind-logs:
	kubectl logs -lapp=comments-api --all-containers=true -f --tail=100

kind-status:
	kubectl get nodes
	kubectl get pods --watch
