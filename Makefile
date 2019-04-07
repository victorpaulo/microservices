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

.PHONY: compose
compose:
	docker-compose up -d

.PHONY: decompose
decompose:
	docker-compose down