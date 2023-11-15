DOCKER_IMAGE_TAG=v$(shell date +%Y%m%d%H%M)

docker-build:
	docker build . -t vortexnotes:$(DOCKER_IMAGE_TAG)

compose-up:
	docker-compose up -d

compose-down:
	docker-compose down
