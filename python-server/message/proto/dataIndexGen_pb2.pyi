from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class EmojiData(_message.Message):
    __slots__ = ("emoji_unicode", "emoji_description")
    EMOJI_UNICODE_FIELD_NUMBER: _ClassVar[int]
    EMOJI_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    emoji_unicode: str
    emoji_description: str
    def __init__(self, emoji_unicode: _Optional[str] = ..., emoji_description: _Optional[str] = ...) -> None: ...

class Action(_message.Message):
    __slots__ = ("action_description", "duration", "start_time", "end_time")
    ACTION_DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    DURATION_FIELD_NUMBER: _ClassVar[int]
    START_TIME_FIELD_NUMBER: _ClassVar[int]
    END_TIME_FIELD_NUMBER: _ClassVar[int]
    action_description: str
    duration: int
    start_time: int
    end_time: int
    def __init__(self, action_description: _Optional[str] = ..., duration: _Optional[int] = ..., start_time: _Optional[int] = ..., end_time: _Optional[int] = ...) -> None: ...

class ActionList(_message.Message):
    __slots__ = ("action_list",)
    ACTION_LIST_FIELD_NUMBER: _ClassVar[int]
    action_list: _containers.RepeatedCompositeFieldContainer[Action]
    def __init__(self, action_list: _Optional[_Iterable[_Union[Action, _Mapping]]] = ...) -> None: ...

class ParsedAction(_message.Message):
    __slots__ = ("emoji_list",)
    EMOJI_LIST_FIELD_NUMBER: _ClassVar[int]
    emoji_list: EmojiData
    def __init__(self, emoji_list: _Optional[_Union[EmojiData, _Mapping]] = ...) -> None: ...
