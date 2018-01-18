.PHONY: build

build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/TheGhostHuCodes/shipper/vessel-service \
		proto/vessel/vessel.proto