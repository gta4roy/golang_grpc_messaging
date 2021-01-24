all:
	compile, build

compile:
	protoc -I order/ order/order.proto --go_out=plugins=grpc:order

build:
	go build -o grpc_messaging/client grpc_messaging/client
	go build -o grpc_messaging/server grpc_messaging/server
	cp -r client/config grpc_messaging/

