@echo off
set out_path=data
python -m grpc_tools.protoc -Idata --python_out=%out_path% --pyi_out=%out_path%  --grpc_python_out=%out_path% data/*.proto