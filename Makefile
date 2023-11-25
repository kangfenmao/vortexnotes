DOCKER_IMAGE_TAG=v$(shell date +%Y%m%d%H%M)
DOCKER_IMAGE_NAME=kangfenmao/vortexnotes:$(DOCKER_IMAGE_TAG)

docker-build:
	docker buildx build --load --platform=linux/arm64 --tag $(DOCKER_IMAGE_NAME) .

docker-push:
	docker buildx build --push --platform=linux/arm64,linux/amd64 --tag $(DOCKER_IMAGE_NAME) .

compose-up:
	docker-compose up -d

compose-down:
	docker-compose down
