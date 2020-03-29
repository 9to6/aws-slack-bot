IMAGE_NAME = aws-slack-bot
IMAGE_VERSION = $(shell cat VERSION)

docker-build:
	@docker build -f docker/Dockerfile -t $(IMAGE_NAME):$(IMAGE_VERSION) .
	@docker tag $(IMAGE_NAME):$(IMAGE_VERSION) $(IMAGE_NAME):latest

docker-run:
	@docker run -d $(IMAGE_NAME) --config="/config.yml"
