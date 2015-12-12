.PHONY: build test start-ircd stop-ircd

build:
	scripts/build.sh

test:
	scripts/test.sh

start-ircd:
	docker run --name=cbx-ircd -d -p 6667:6667 xena/elemental-ircd
	@echo Development IRCD running at localhost:6667
	@echo To stop the ircd, run make stop-ircd

stop-ircd:
	docker rm -f cbx-ircd
