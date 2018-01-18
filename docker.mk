#
# ---------
# docker.mk
# ---------
#
# Makefile targets useful for running dockerized services.
#
DOCKER_NW       ?= $(APPNAME)_nw
DOCKER_RUN_OPTS += --network $(DOCKER_NW)

docker_nw: .docker_nw
.docker_nw:
	@echo "Creating $(DOCKER_NW) network..."
	@docker network create -d bridge $(DOCKER_NW) > /dev/null
	@echo "Created $(DOCKER_NW) network."
	@touch $@

docker_clean:
	@-docker network rm $(DOCKER_NW) > /dev/null
	@rm -f .docker_nw

.PHONY: docker_clean
