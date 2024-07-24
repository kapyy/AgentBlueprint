import signal
import sys
import threading
import time
from concurrent import futures
import logging
import grpc

from message.proto.pythonServicer_pb2_grpc import APMServiceServicer, add_APMServiceServicer_to_server
from message.proto.pythonServicer_pb2 import ServiceResponse
from factory.factory_handler import DeserializeFunctionNode, CallSubordinateFunction
from util.service_config import ServiceConfig

logger = logging.getLogger("kotoko.debug")

config = ServiceConfig()
config.embedding_model = 'all-MiniLM-L6-v2'
config.host = '192.168.50.43'
config.port = '19530'
config.embedding_device = 'cuda'


class BlueprintServicer(APMServiceServicer):
    def MainServiceRequest(self, request, context):
        # logger.debug("MainServiceRequest: %s", str(request))
        result, resultstr = DeserializeFunctionNode(request.data)
        return ServiceResponse(message_id=request.data.node_id, res_data=result.SerializeToString())

    def SubordinateServiceRequest(self, request, context):
        # logger.debug("Subordinate: %s", str(request))
        result = CallSubordinateFunction(request.message_id, request.data_type, request.rqst_data)

        return ServiceResponse(message_id=request.message_id, res_data=result.SerializeToString())


# class FunctionalServicer(SubFunctionalServiceServicer):
#
# TODO Need to implement Embedding functions without towhee pipeline
# def EmbeddingNounChunks(self, request, context):
#     logger.debug("EmbeddedNounChunk: %s", str(request.prompt_sentence))
#
#     pres = ParseNounChunks(request.prompt_sentence)
#
#     logger.debug("result: %s", str(pres))
#     q = pipes["embedding_vec_pipe"]
#     word_list = WordList()
#     for index, word in enumerate(pres):
#         word_vec = WordVec()
#         word_vec.word = word
#         word_vec.dimension = 384
#         vec = q(word).get()[0]
#         # logger.debug("vec: %s", str(vec))
#         word_vec.vec.extend(vec)
#         word_list.words.append(word_vec)
#     return word_list
#
# def EmbeddingSentence(self, request, context):
#     logger.debug("EmbeddedSentence: %s", str(request.prompt_sentence))
#     q = pipes["embedding_vec_pipe"]
#     vec = q(request.prompt_sentence).get()[0]
#     result = SentenceVec()
#     result.dimension = 384
#     result.sentence = request.prompt_sentence
#     result.vec.extend(vec)
#
#     return result
# def EmbeddingList(self, request, context):
#     logger.debug("Embedding List: %s", str(request.list))
#     # log senders
#     q = pipes["embedding_vec_pipe"]
#     input_list = request.list
#     vecs = q.batch(input_list)
#     word_list = WordList()
#     for index, word in enumerate(vecs):
#         word_vec = WordVec()
#         word_vec.word = input_list[index]
#         word_vec.dimension = 384
#         vec = word.get()[0]
#         # logger.debug("vec: %s", str(vec))
#         word_vec.vec.extend(vec)
#         word_list.words.append(word_vec)
#     return word_list

class CoreServerThread(threading.Thread):
    def __init__(self):
        threading.Thread.__init__(self)
        self.port = "50051"
        self.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        self.stop_event = threading.Event()

    def run(self):
        add_APMServiceServicer_to_server(BlueprintServicer(), self.server)
        # add_SubFunctionalServiceServicer_to_server(FunctionalServicer(), server)
        self.server.add_insecure_port("[::]:" + self.port)
        # debug grpc
        # logger.info("Server started, listening on " + self.port)
        print("Core Server started, listening on " + self.port)
        self.server.start()
        self.stop_event.wait()
        self.server.stop(0)

    def stop(self):
        self.stop_event.set()

