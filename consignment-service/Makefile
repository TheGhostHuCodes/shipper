.PHONY: build clean

build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/TheGhostHuCodes/shipper/consignment-service \
		proto/consignment/consignment.proto

clean:
	rm proto/*/*.pb.go
