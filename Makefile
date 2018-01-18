APPNAME = self

include docker.mk
include nomad.mk

start: nomad_start

clean: nomad_clean consul_clean
