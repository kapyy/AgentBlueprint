import configparser

import grpc
import logging

import grpc

import message.data.dataIndexGen_pb2
from message.data.functionDistribute_pb2 import GeneralPyRequest
from message.data.functionDistribute_pb2_grpc import APMFunctionsServiceStub
import apmFunctionFactory.apm_factory_data as apm_factory_data

config_parser = configparser.ConfigParser()
config_parser.read('./conf.ini')

def SerializeDataWithTypeID(data, data_id):
    type_id = data_id
    if type_id == 0:
        return "No Data Type"
    elif type_id == 1007:
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


# currently only max of 2 inputs are accepted


METHOD_REFLECTION = {
    1100100100: "GetDayPlan",
    1100100200: "GetInBetweenPlans",
    1100100300: "ExecuteActionWithPlans",
    1300100100: "ExecuteActionWithWhisper",
    1100100400: "MemoryDistillToLongTerm",
    1100100500: "SummarizeAgent",
    1100100600: "ExecuteActionWithObservation",
    1100100700: "ObservationIntoMemoryLog",
    1210100100: "MemoryScoring",
    1210100200: "ActionFormatter",
    1210100300: "ObjectFormatter",
    1200200100: "ObservationReaction",
    1300300100: "GetChatResponse",
    1300400100: "GetTalkContent",
}
