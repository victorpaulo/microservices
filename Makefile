
.PHONY: build
build:
	cd app1 && make build
	cd app2 && make build

.PHONY: remove
remove:
	cd app1 && make remove
	cd app2 && make remove

.PHONY: run
run:
	cd app1 && make run
	cd app2 && make run

.PHONY: stop
stop:
	cd app1 && make stop
	cd app2 && make stop

.PHONY: push
push:
	cd app1 && make push
	cd app2 && make push


.PHONY: compose
compose:
	docker-compose up -d

.PHONY: decompose
decompose:
	docker-compose down

.PHONY: docker-list
docker-list:
	curl -X GET http://172.28.128.3:5000/v2/_catalog

k8s-install-apps:
	cd app1 && make k8s-install
	cd app2 && make k8s-install

delele-registry-images:
	cd app1 && make del-docker-registry
	cd app2 && make del-docker-registry