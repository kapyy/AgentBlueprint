import configparser
import logging
import grpc
import message.data.dataIndexGen_pb2
from message.data.functionDistribute_pb2 import GeneralPyRequest
from message.data.functionDistribute_pb2_grpc import APMFunctionsServiceStub

config_parser = configparser.ConfigParser()
config_parser.read('./conf.ini')

#Because all the Mainservice pipeline request's Data are formatted before it goes into factory
#So only Static Function's input data need to be reflected
def SerializeDataWithTypeID(data, data_id):
    if data_id == 0:
        return "No Data Type"
    elif data_id == 1007:
        action = message.data.dataIndexGen_pb2.Action()
        action.ParseFromString(data)
        return action
    else:
        return None


def SubServicerDistributionCaller(id, rqst):
    with grpc.insecure_channel(config_parser['server.setting']['FunctionURL']) as channel:
        stub = APMFunctionsServiceStub(channel)
        if id == 400100100:
            return stub.ActionFormatter(rqst)


#The Second Value are returned when Minor Function is used only
def MainServicerDistributorCaller(id, context):
    with grpc.insecure_channel(config_parser['server.setting']['FunctionURL']) as channel:
        try:
            grpc.channel_ready_future(channel).result(timeout=10)
        except grpc.FutureTimeoutError:
            logging.error("Error connecting to py function server at %s", config_parser['server.setting']['FunctionURL'])
            return None, "No Data Type"
        stub = APMFunctionsServiceStub(channel)

        if id == 100100100:
            request = GeneralPyRequest(prompt=context["prompt"])
            return stub.GetDayPlan(request), "No String Format"

        elif id == 100200400:
            request = GeneralPyRequest(prompt=context["prompt"], system_prompt=context["system"])
            return stub.GenerateFatigueHabitAction(request), "No String Format"
        elif id == 310400400:
            request = GeneralPyRequest(prompt=context["prompt"], text=context["text_input"])
            return stub.InsertActionsWithObservation(request), "No String Format"

        else:
            print("No such service Function found")
            return None, "No Data Type"

