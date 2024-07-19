import dataIndexGen_pb2 as _dataIndexGen_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class GeneralPyRequest(_message.Message):
    __slots__ = ("prompt", "text", "system_prompt")
    PROMPT_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    SYSTEM_PROMPT_FIELD_NUMBER: _ClassVar[int]
    prompt: str
    text: str
    system_prompt: str
    def __init__(self, prompt: _Optional[str] = ..., text: _Optional[str] = ..., system_prompt: _Optional[str] = ...) -> None: ...
