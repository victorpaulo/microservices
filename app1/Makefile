APP=microservices_app1
VERSION=1.1
PORT=8081
DANGLING_IMAGES_CMD=$(shell docker images -q -f dangling=true)
DOCKER_IMAGES=$(shell docker images | grep $(APP) | tr -s ' ' | cut -d ' ' -f 3)
DOCKER_REGISTRY=172.28.128.5:5000
KUBECONFIG_PATH=/Users/victorp/projects/kubernates-vm
CURL_DIGEST=$(shell curl -v --silent -H "Accept: application/vnd.docker.distribution.manifest.v2+json" \
		-X GET http://$(DOCKER_REGISTRY)/v2/$(APP)/manifests/$(VERSION) 2>&1 \
		| grep Docker-Content-Digest | awk '{print ($3)}')


.PHONY: remove
remove: stop
	-docker rmi --force $(DANGLING_IMAGES_CMD)
	-docker rmi --force $(DOCKER_IMAGES)
	

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
	#echo "$(CURL_DIGEST)"
	curl -v --silent -H "Accept: application/vnd.docker.distribution.manifest.v2+json" \
		-X DELETE http://$(DOCKER_REGISTRY)/v2/$(APP)/manifests/$(CURL_DIGEST)