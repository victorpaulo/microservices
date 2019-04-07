APP=microservices_app2
VERSION=1.0
PORT=8082
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
	kubectl apply -f k8s/deployment.yaml
	kubectl apply -f k8s/service.yaml

del-docker-registry:
	#get digest id
	curl -v --silent -H "Accept: application/vnd.docker.distribution.manifest.v2+json" \
		-X GET http://$(DOCKER_REGISTRY)/v2/$(APP)/manifests/$(VERSION) 2>&1 \
		| grep Docker-Content-Digest | awk '{print ($3)}'
	#delete from registry
	curl -v --silent -H "Accept: application/vnd.docker.distribution.manifest.v2+json" \
		-X DELETE http://$(DOCKER_REGISTRY)/v2/$(APP)/manifests/$(digest_id)