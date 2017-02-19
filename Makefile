.PHONY: test

PACKAGES=`go list ./... | grep -v /vendor/`

test:
	go test -v ${PACKAGES}
