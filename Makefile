COLOR := $(shell git rev-parse HEAD | cut -c 1-6)
DOCKER_IMAGE ?= mnemonic-buffalo


.PHONY: build
build: mnemonic-buffalo

mnemonic-buffalo: main.go
	buffalo build -o mnemonic-buffalo-$(COLOR)

.PHONY: image
image:
	@echo Building with color $(COLOR)
	COLOR=$(COLOR) docker build . -t alexlovelltroy/$(DOCKER_IMAGE):$(COLOR) --build-arg COLOR=$(COLOR) -f Dockerfile

.PHONY: push
push: image
	docker push alexlovelltroy/$(DOCKER_IMAGE):$(COLOR)


.PHONY: run
run:
	docker run -it -p "3000:3000" alexlovelltroy/$(DOCKER_IMAGE):$(COLOR)
