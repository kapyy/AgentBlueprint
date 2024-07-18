@echo off
set out_path=protoData
protoc --proto_path=protoData --go_out=%out_path% --go_opt=module=golang-client/message --go-grpc_out=%out_path% --go-grpc_opt=module=golang-client/message protoData/*.proto