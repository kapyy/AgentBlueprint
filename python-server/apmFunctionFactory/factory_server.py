"""The reflection-enabled version of gRPC helloworld.Greeter server."""

from concurrent import futures
import logging

import grpc
from grpc_reflection.v1alpha import reflection
import message.data.pyFunctionDistribute_pb2_grpc
import message.data.pyFunctionDistribute_pb2
from core.function_servicer_implementation import APMFunctionsServiceServicer


def FactoryServerStart():
    port = "50077"
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    apm.pyFunctionDistribute_pb2_grpc.add_APMFunctionsServiceServicer_to_server(
        APMFunctionsServiceServicer(), server)
    SERVER_NAMES = (
        apm.pyFunctionDistribute_pb2.DESCRIPTOR.services_by_name['APMFunctionsService'].full_name,
        reflection.SERVICE_NAME,
    )

    reflection.enable_server_reflection(SERVER_NAMES, server)
    server.add_insecure_port("[::]:" + port)
    server.start()
    logging.getLogger().info("Server started, listening on " + port)
    server.wait_for_termination()
