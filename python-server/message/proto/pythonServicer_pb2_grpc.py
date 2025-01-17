# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from message.proto import pythonServicer_pb2 as message_dot_proto_dot_pythonServicer__pb2

GRPC_GENERATED_VERSION = '1.65.1'
GRPC_VERSION = grpc.__version__
EXPECTED_ERROR_RELEASE = '1.66.0'
SCHEDULED_RELEASE_DATE = 'August 6, 2024'
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    warnings.warn(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in message/proto/pythonServicer_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class APMServiceStub(object):
    """Main Service

    We removed single module rpc calls from .proto,
    which mean we will execute every function module within only python server,
    we can use a single rpc call to send all the data Python server need
    and let Python server to deserialize .apm file to fulfill the request.

    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.MainServiceRequest = channel.unary_unary(
                '/protoData.APMService/MainServiceRequest',
                request_serializer=message_dot_proto_dot_pythonServicer__pb2.MainServicerRequest.SerializeToString,
                response_deserializer=message_dot_proto_dot_pythonServicer__pb2.ServiceResponse.FromString,
                _registered_method=True)
        self.SubordinateServiceRequest = channel.unary_unary(
                '/protoData.APMService/SubordinateServiceRequest',
                request_serializer=message_dot_proto_dot_pythonServicer__pb2.SubordinateServicerRequest.SerializeToString,
                response_deserializer=message_dot_proto_dot_pythonServicer__pb2.ServiceResponse.FromString,
                _registered_method=True)


class APMServiceServicer(object):
    """Main Service

    We removed single module rpc calls from .proto,
    which mean we will execute every function module within only python server,
    we can use a single rpc call to send all the data Python server need
    and let Python server to deserialize .apm file to fulfill the request.

    """

    def MainServiceRequest(self, request, context):
        """Upload MainServiceStructure to Python Server
        returned with pre-defined data structure and its formatted data for game use
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SubordinateServiceRequest(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_APMServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'MainServiceRequest': grpc.unary_unary_rpc_method_handler(
                    servicer.MainServiceRequest,
                    request_deserializer=message_dot_proto_dot_pythonServicer__pb2.MainServicerRequest.FromString,
                    response_serializer=message_dot_proto_dot_pythonServicer__pb2.ServiceResponse.SerializeToString,
            ),
            'SubordinateServiceRequest': grpc.unary_unary_rpc_method_handler(
                    servicer.SubordinateServiceRequest,
                    request_deserializer=message_dot_proto_dot_pythonServicer__pb2.SubordinateServicerRequest.FromString,
                    response_serializer=message_dot_proto_dot_pythonServicer__pb2.ServiceResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'protoData.APMService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('protoData.APMService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class APMService(object):
    """Main Service

    We removed single module rpc calls from .proto,
    which mean we will execute every function module within only python server,
    we can use a single rpc call to send all the data Python server need
    and let Python server to deserialize .apm file to fulfill the request.

    """

    @staticmethod
    def MainServiceRequest(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/protoData.APMService/MainServiceRequest',
            message_dot_proto_dot_pythonServicer__pb2.MainServicerRequest.SerializeToString,
            message_dot_proto_dot_pythonServicer__pb2.ServiceResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def SubordinateServiceRequest(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/protoData.APMService/SubordinateServiceRequest',
            message_dot_proto_dot_pythonServicer__pb2.SubordinateServicerRequest.SerializeToString,
            message_dot_proto_dot_pythonServicer__pb2.ServiceResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)


class SubFunctionalServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.EmbeddingNounChunks = channel.unary_unary(
                '/protoData.SubFunctionalService/EmbeddingNounChunks',
                request_serializer=message_dot_proto_dot_pythonServicer__pb2.RequestPrompt.SerializeToString,
                response_deserializer=message_dot_proto_dot_pythonServicer__pb2.WordList.FromString,
                _registered_method=True)
        self.EmbeddingSentence = channel.unary_unary(
                '/protoData.SubFunctionalService/EmbeddingSentence',
                request_serializer=message_dot_proto_dot_pythonServicer__pb2.RequestPrompt.SerializeToString,
                response_deserializer=message_dot_proto_dot_pythonServicer__pb2.SentenceVec.FromString,
                _registered_method=True)
        self.EmbeddingList = channel.unary_unary(
                '/protoData.SubFunctionalService/EmbeddingList',
                request_serializer=message_dot_proto_dot_pythonServicer__pb2.RequestList.SerializeToString,
                response_deserializer=message_dot_proto_dot_pythonServicer__pb2.WordList.FromString,
                _registered_method=True)
        self.EmbeddingTopic = channel.unary_unary(
                '/protoData.SubFunctionalService/EmbeddingTopic',
                request_serializer=message_dot_proto_dot_pythonServicer__pb2.RequestPrompt.SerializeToString,
                response_deserializer=message_dot_proto_dot_pythonServicer__pb2.WordList.FromString,
                _registered_method=True)


class SubFunctionalServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def EmbeddingNounChunks(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def EmbeddingSentence(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def EmbeddingList(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def EmbeddingTopic(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SubFunctionalServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'EmbeddingNounChunks': grpc.unary_unary_rpc_method_handler(
                    servicer.EmbeddingNounChunks,
                    request_deserializer=message_dot_proto_dot_pythonServicer__pb2.RequestPrompt.FromString,
                    response_serializer=message_dot_proto_dot_pythonServicer__pb2.WordList.SerializeToString,
            ),
            'EmbeddingSentence': grpc.unary_unary_rpc_method_handler(
                    servicer.EmbeddingSentence,
                    request_deserializer=message_dot_proto_dot_pythonServicer__pb2.RequestPrompt.FromString,
                    response_serializer=message_dot_proto_dot_pythonServicer__pb2.SentenceVec.SerializeToString,
            ),
            'EmbeddingList': grpc.unary_unary_rpc_method_handler(
                    servicer.EmbeddingList,
                    request_deserializer=message_dot_proto_dot_pythonServicer__pb2.RequestList.FromString,
                    response_serializer=message_dot_proto_dot_pythonServicer__pb2.WordList.SerializeToString,
            ),
            'EmbeddingTopic': grpc.unary_unary_rpc_method_handler(
                    servicer.EmbeddingTopic,
                    request_deserializer=message_dot_proto_dot_pythonServicer__pb2.RequestPrompt.FromString,
                    response_serializer=message_dot_proto_dot_pythonServicer__pb2.WordList.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'protoData.SubFunctionalService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('protoData.SubFunctionalService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class SubFunctionalService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def EmbeddingNounChunks(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/protoData.SubFunctionalService/EmbeddingNounChunks',
            message_dot_proto_dot_pythonServicer__pb2.RequestPrompt.SerializeToString,
            message_dot_proto_dot_pythonServicer__pb2.WordList.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def EmbeddingSentence(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/protoData.SubFunctionalService/EmbeddingSentence',
            message_dot_proto_dot_pythonServicer__pb2.RequestPrompt.SerializeToString,
            message_dot_proto_dot_pythonServicer__pb2.SentenceVec.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def EmbeddingList(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/protoData.SubFunctionalService/EmbeddingList',
            message_dot_proto_dot_pythonServicer__pb2.RequestList.SerializeToString,
            message_dot_proto_dot_pythonServicer__pb2.WordList.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def EmbeddingTopic(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/protoData.SubFunctionalService/EmbeddingTopic',
            message_dot_proto_dot_pythonServicer__pb2.RequestPrompt.SerializeToString,
            message_dot_proto_dot_pythonServicer__pb2.WordList.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
