@echo off

set py_out_path=python-server
set go_module=golang-client/message
set go_out_path=golang-client/message/proto
set proto_path=proto_message
set src_path=proto_message/message/proto
protoc -I%proto_path% --go_out=%go_out_path% --go_opt=module=%go_module% --go-grpc_out=%go_out_path% --go-grpc_opt=module=%go_module% %src_path%/*.proto
python -m grpc_tools.protoc -I%proto_path% --python_out=%py_out_path% --pyi_out=%py_out_path%  --grpc_python_out=%py_out_path% %src_path%/*.proto
