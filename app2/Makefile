APP=app2
VERSION=1.0
PORT=8082

.PHONY: remove
remove:
	docker rmi $(APP):$(VERSION)

.PHONY: build
build:
	docker build -t $(APP):$(VERSION) .

.PHONY: run
run:
	docker run --name $(APP) --rm -d -p $(PORT):$(PORT) $(APP):$(VERSION)

.PHONY: stop
stop:
	docker stop $(APP)