APP=app2
VERSION=1.0
PORT=8082

remove:
	docker rmi $(APP):$(VERSION)

build:
	docker build -t $(APP):$(VERSION) .

run:
	docker run --name $(APP) --rm -d -p $(PORT):$(PORT) $(APP):$(VERSION)

stop:
	docker stop $(APP)