APPNAME              = selftest
IMG                 ?= bsorahan/
BUILD_IMG            = $(APPNAME):build
VERSION             ?= latest
RELEASE_IMG          = $(APPNAME):$(VERSION)
SRCS                 = $(wildcard *.go)
EXTRA_ASSETS        ?= /etc/ssl

clean:
	@rm -rf .image .assets.tar

release: .release

.release: Dockerfile.release .assets.tar
	docker build --rm -t $(RELEASE_IMG) -f Dockerfile.release .
	@touch $@

push: .release
	docker push $(RELEASE_IMG)

release: .assets.tar

.assets.tar: Dockerfile $(SRCS)
	docker build --pull --rm --build-arg APPNAME=$(APPNAME) -t $(BUILD_IMG) .
	docker run --rm $(BUILD_IMG) tar cf - /app $(EXTRA_ASSETS) > $@

.PHONY: clean push release
