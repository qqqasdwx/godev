IMAGE_NAME=registry.cn-beijing.aliyuncs.com/$YOUR_PROJECT_PATH

all:
	docker build -t $(IMAGE_NAME) .
push:
	docker push $(IMAGE_NAME)
clean:
	docker rmi  $(IMAGE_NAME)