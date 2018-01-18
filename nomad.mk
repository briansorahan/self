#
# --------
# nomad.mk
# --------
#
# Makefile targets for services that use nomad.
#
# Variables:
#     CONSUL_VERSION       Consul version.
#     NOMAD_VERSION        Nomad version.
#     NOMAD_I              Docker image to use for nomad. Defaults to djenriquez/nomad.
#     CONSUL_I             Docker image to use for consul. Defaults to library/consul.
#
CONSUL_VERSION     ?= 0.9.3
NOMAD_VERSION      ?= v0.6.3
NOMAD_I            ?= djenriquez/nomad:$(NOMAD_VERSION)
CONSUL_I           ?= consul:$(CONSUL_VERSION)
CONSUL_C           ?= $(APPNAME)_consul
NOMADCLIENT_C      ?= $(APPNAME)_nomadclient
NOMADSERVER_C      ?= $(APPNAME)_nomadserver
NOMADCLIENT_CONFIG ?='{"client":{"enabled":true},"consul":{"address":"consul:8500"},"bind_addr":"0.0.0.0","enable_debug":true}'
NOMADSERVER_CONFIG ?='{"server":{"enabled":true,"bootstrap_expect":1},"bind_addr":"0.0.0.0","consul":{"address":"consul:8500"},"enable_debug":true}'
NOMAD_PORT         ?= 4646
CONSUL_PORT        ?= 8500
NOMAD_VOLUMES      ?= -v /tmp:/tmp -v /var/run/docker.sock:/var/run/docker.sock

nomad_start: .consul_start .nomad_server .nomad_client

.nomad_server: .docker_nw
	@echo "Running nomad server..."
	@docker run -d \
           --privileged \
           --link $(CONSUL_C):consul \
           --network $(DOCKER_NW) \
           --name $(NOMADSERVER_C) \
           -p $(NOMAD_PORT) \
           -e NOMAD_LOCAL_CONFIG=$(NOMADSERVER_CONFIG) \
           $(NOMAD_VOLUMES) \
           $(NOMAD_I) agent >/dev/null
	@echo "Nomad server up and running."
	@touch $@

.nomad_client: .docker_nw
	@echo "Running nomad client..."
	@docker run -d \
           --privileged \
           --link $(CONSUL_C):consul \
           --network $(DOCKER_NW) \
           --name $(NOMADCLIENT_C) \
           -p $(NOMAD_PORT) \
           -e NOMAD_LOCAL_CONFIG=$(NOMADCLIENT_CONFIG) \
           $(NOMAD_VOLUMES) \
           $(NOMAD_I) agent >/dev/null
	@echo "Nomad client up and running."
	@touch $@

.consul_start: .docker_nw
	@echo "Running consul..."
	@docker run -d -p $(CONSUL_PORT) --network $(DOCKER_NW) --name $(CONSUL_C) $(CONSUL_I) >/dev/null
	@echo "consul up and running."
	@touch $@

consul_clean:
	@echo "Stopping consul..."
	@-docker rm -f $(CONSUL_C) 2>&1 | grep -v "No such container"
	@echo "Stopped consul."
	@-rm -rf .consul_*

nomad_clean:
	@echo "Stopping nomad cluster..."
	@-docker rm -f $(NOMADSERVER_C) $(NOMADCLIENT_C) 2>&1 | grep -v "No such container"
	@echo "Stopped nomad cluster."
	@-rm -rf .nomad_*

.PHONY: consul_clean nomad_clean nomad_start
