
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
    elif data_id == 1001:
        action = message.proto.dataIndexGen_pb2.Action()
        action.ParseFromString(data)
        return action
    elif data_id == 4001:
        parsedaction = message.proto.dataIndexGen_pb2.ParsedAction()
        parsedaction.ParseFromString(data)
        return parsedaction
    else:
        return None

def StaticFunctionDistributionCaller(id, rqst):
    with grpc.insecure_channel(config_parser['server.setting']['FunctionURL']) as channel:
        stub = APMFunctionsServiceStub(channel)
        if id == 10000002:
            return stub.ActionFormatter(rqst)

def MainServicerDistributorCaller(id, context):
    with grpc.insecure_channel(config_parser['server.setting']['FunctionURL']) as channel:
        try:
            grpc.channel_ready_future(channel).result(timeout=10)
        except grpc.FutureTimeoutError:
            logging.error("Error connecting to py function server at %s", config_parser['server.setting']['FunctionURL'])
            return None, "No Data Type"
        stub = APMFunctionsServiceStub(channel)
        if id == 10000001:
            request = GeneralPyRequest(prompt=context["prompt"],text=context["text_input"],system_prompt=context["system"])
            return stub.InsertActionWithObservation(request), "No Data String Out needed"
        else:
            logging.error("No such service Function found")
            return None, "No Data Type"
