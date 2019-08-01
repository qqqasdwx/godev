IMAGE_PREFIX=registry.cn-beijing.aliyuncs.com/intwallyun
IMAGE_NAME=$(IMAGE_PREFIX)/{{.ProjectName}}
DOCKER_TAG_NAME=2.0
DOCKERFILE_PATH=./Dockerfile
all:
	docker build -f $(DOCKERFILE_PATH) -t $(IMAGE_NAME):$(DOCKER_TAG_NAME) .
push:
	docker push $(IMAGE_NAME):$(DOCKER_TAG_NAME)
clean:
	docker rmi  $(IMAGE_NAME):$(DOCKER_TAG_NAME)