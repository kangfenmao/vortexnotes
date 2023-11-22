DOCKER_IMAGE_TAG=v$(shell date +%Y%m%d%H%M)

docker-build:
	docker buildx build --load --platform=linux/arm64 --tag vortexnotes:latest .

compose-up:
	docker-compose up -d

compose-down:
	docker-compose down
