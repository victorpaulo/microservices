APP=microservices_app1
VERSION=1.0
PORT=8081
DANGLING_IMAGES_CMD=$(shell docker images -q -f dangling=true)
DOCKER_REGISTRY=172.28.128.5:5000
KUBECONFIG_PATH=/Users/victorp/projects/kubernates-vm

.PHONY: remove
remove: stop
	-docker rmi $(DANGLING_IMAGES_CMD)
	docker rmi $(APP):$(VERSION)

.PHONY: build
build:
	docker build -t $(APP):$(VERSION) .
	docker tag $(APP):$(VERSION) $(DOCKER_REGISTRY)/$(APP):$(VERSION)

.PHONY: run
run:
	docker run --name $(APP) --rm -d -p $(PORT):$(PORT) $(APP):$(VERSION)

.PHONY: stop
stop:
	echo "stopping"
	docker stop $(APP)

push:
	docker push $(DOCKER_REGISTRY)/$(APP):$(VERSION)

k8s-local:
	export KUBECONFIG=$(KUBECONFIG_PATH)/config
	kubectl cluster-info

k8s-install: k8s-local
	kubectl create -f k8s/deployment.yaml
	kubectl create -f k8s/service.yaml