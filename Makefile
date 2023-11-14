DOCKER_IMAGE_TAG=v$(shell date +%Y%m%d%H%M)

build_docker:
	docker build . -f docker/Dockerfile -t vortexnotes:$(DOCKER_IMAGE_TAG)

compose_up:
	docker-compose -f docker/docker-compose.yml up -d
