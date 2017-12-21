TAG?=latest
SQUASH?=false

build:
	docker build --build-arg http_proxy="${http_proxy}" --build-arg https_proxy="${https_proxy}" -t functions/faas-swarm:$(TAG) . --squash=${SQUASH}

push:
	docker push functions/faas-swarm:$(TAG)

all: build