from message.proto import APMFactory_pb2 as _APMFactory_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class MainServicerRequest(_message.Message):
    __slots__ = ("message_id", "data")
    MESSAGE_ID_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    message_id: int
    data: _APMFactory_pb2.NodeData
    def __init__(self, message_id: _Optional[int] = ..., data: _Optional[_Union[_APMFactory_pb2.NodeData, _Mapping]] = ...) -> None: ...

class SubordinateServicerRequest(_message.Message):
    __slots__ = ("message_id", "data_type", "rqst_data")
    MESSAGE_ID_FIELD_NUMBER: _ClassVar[int]
    DATA_TYPE_FIELD_NUMBER: _ClassVar[int]
    RQST_DATA_FIELD_NUMBER: _ClassVar[int]
    message_id: int
    data_type: int
    rqst_data: bytes
    def __init__(self, message_id: _Optional[int] = ..., data_type: _Optional[int] = ..., rqst_data: _Optional[bytes] = ...) -> None: ...

class ServiceResponse(_message.Message):
    __slots__ = ("message_id", "res_data")
    MESSAGE_ID_FIELD_NUMBER: _ClassVar[int]
    RES_DATA_FIELD_NUMBER: _ClassVar[int]
    message_id: int
    res_data: bytes
    def __init__(self, message_id: _Optional[int] = ..., res_data: _Optional[bytes] = ...) -> None: ...

class RequestPrompt(_message.Message):
    __slots__ = ("prompt_sentence",)
    PROMPT_SENTENCE_FIELD_NUMBER: _ClassVar[int]
    prompt_sentence: str
    def __init__(self, prompt_sentence: _Optional[str] = ...) -> None: ...

class RequestList(_message.Message):
    __slots__ = ("list",)
    LIST_FIELD_NUMBER: _ClassVar[int]
    list: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, list: _Optional[_Iterable[str]] = ...) -> None: ...

class WordList(_message.Message):
    __slots__ = ("words",)
    WORDS_FIELD_NUMBER: _ClassVar[int]
    words: _containers.RepeatedCompositeFieldContainer[WordVec]
    def __init__(self, words: _Optional[_Iterable[_Union[WordVec, _Mapping]]] = ...) -> None: ...

class WordVec(_message.Message):
    __slots__ = ("word", "dimension", "vec")
    WORD_FIELD_NUMBER: _ClassVar[int]
    DIMENSION_FIELD_NUMBER: _ClassVar[int]
    VEC_FIELD_NUMBER: _ClassVar[int]
    word: str
    dimension: int
    vec: _containers.RepeatedScalarFieldContainer[float]
    def __init__(self, word: _Optional[str] = ..., dimension: _Optional[int] = ..., vec: _Optional[_Iterable[float]] = ...) -> None: ...

class SentenceVec(_message.Message):
    __slots__ = ("sentence", "dimension", "vec")
    SENTENCE_FIELD_NUMBER: _ClassVar[int]
    DIMENSION_FIELD_NUMBER: _ClassVar[int]
    VEC_FIELD_NUMBER: _ClassVar[int]
    sentence: str
    dimension: int
    vec: _containers.RepeatedScalarFieldContainer[float]
    def __init__(self, sentence: _Optional[str] = ..., dimension: _Optional[int] = ..., vec: _Optional[_Iterable[float]] = ...) -> None: ...
