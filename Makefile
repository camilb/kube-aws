CALICO_BUILD?=calico/go-build
PACKAGE_NAME?=kubernetes-incubator/kube-aws
LOCAL_USER_ID?=$(shell id -u $$USER)

.PHONY: build
build:
	./build

vendor: glide.yaml
	rm -f glide.lock
	docker run --rm \
	    -v $(CURDIR):/go/src/github.com/$(PACKAGE_NAME):rw \
	    -e LOCAL_USER_ID=$(LOCAL_USER_ID) \
            -w /go/src/github.com/$(PACKAGE_NAME) \
	    $(CALICO_BUILD) glide install -strip-vendor
	

.PHONY: format
format:
	test -z "$$(find . -path ./vendor -prune -type f -o -name '*.go' -exec gofmt -d {} + | tee /dev/stderr)" || \
	test -z "$$(find . -path ./vendor -prune -type f -o -name '*.go' -exec gofmt -w {} + | tee /dev/stderr)"

.PHONY: test
test: build
	./make/test

.PHONY: test-with-cover
test-with-cover: build
	./make/test with-cover
