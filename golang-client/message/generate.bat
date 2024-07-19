@echo off
set out_path=message/protoData
set proto_path=message/protoData
protoc --proto_path=%proto_path% --go_out=%out_path% --go_opt=module=golang-client/message --go-grpc_out=%out_path% --go-grpc_opt=module=golang-client/message %proto_path%/*.proto