DOCKER_IMAGE_TAG=v$(shell date +%Y%m%d%H%M)

build_docker:
	docker build . -t vortexnotes:$(DOCKER_IMAGE_TAG)

compose_up:
	docker-compose up -d
