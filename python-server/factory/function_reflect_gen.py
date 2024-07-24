
import configparser
import logging
import grpc
import message.proto.dataIndexGen_pb2
from message.proto.functionDistribute_pb2 import GeneralPyRequest
from message.proto.functionDistribute_pb2_grpc import APMFunctionsServiceStub
config_parser = configparser.ConfigParser()
config_parser.read('./conf.ini')

def SerializeDataWithTypeID(data, data_id):
    if data_id == 0:
        return "No Data Type"
