.PHONY: build

build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/TheGhostHuCodes/shipper/consignment-service \
		proto/consignment/consignment.proto