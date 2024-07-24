"""The reflection-enabled version of gRPC helloworld.Greeter server."""
import signal
import sys
import threading
import time
from concurrent import futures
import logging

import grpc
from grpc_reflection.v1alpha import reflection

from message.proto import functionDistribute_pb2,functionDistribute_pb2_grpc

from function.function_servicer_implementation import APMFunctionsServiceServicer

stop_event = threading.Event()


class FunctionServerThread(threading.Thread):
    def __init__(self):
        threading.Thread.__init__(self)
        self.port = "50077"
        self.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        self.stop_event = threading.Event()

    def run(self):
        functionDistribute_pb2_grpc.add_APMFunctionsServiceServicer_to_server(
            APMFunctionsServiceServicer(), self.server)
        SERVER_NAMES = (
            functionDistribute_pb2.DESCRIPTOR.services_by_name['APMFunctionsService'].full_name,
            reflection.SERVICE_NAME,
        )

        reflection.enable_server_reflection(SERVER_NAMES, self.server)
        self.server.add_insecure_port("[::]:" + self.port)
        self.server.start()
        print("Function Server started, listening on " + self.port)
        self.stop_event.wait()
        self.server.stop(0)

    def stop(self):
        self.stop_event.set()