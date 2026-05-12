DOCKER_COMMAND=docker
IMAGE_NAME=ieeexplore-mcp
CONTAINER_NAME=ieeexplore-mcp

.PHONY: build run stop rm

build:
	$(DOCKER_COMMAND) build -t $(IMAGE_NAME) .

run: build
	$(DOCKER_COMMAND) run -p 8080:8080 --name $(CONTAINER_NAME) -d $(IMAGE_NAME)

stop:
	$(DOCKER_COMMAND) stop $(CONTAINER_NAME)

rm:
	$(DOCKER_COMMAND) rm $(CONTAINER_NAME)

