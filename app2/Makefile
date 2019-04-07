APP=microservices_app2
VERSION=1.0
PORT=8082
DANGLING_IMAGES_CMD=$(shell docker images -q -f dangling=true)

.PHONY: remove
remove: stop
	-docker rmi $(DANGLING_IMAGES_CMD)
	docker rmi $(APP):$(VERSION)

.PHONY: build
build:
	docker build -t $(APP):$(VERSION) .

.PHONY: run
run:
	docker run --name $(APP) --rm -d -p $(PORT):$(PORT) $(APP):$(VERSION)

.PHONY: stop
stop:
	echo "stopping"
	docker stop $(APP)