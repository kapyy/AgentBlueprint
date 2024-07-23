from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class NodeConnector(_message.Message):
    __slots__ = ("input_nodes",)
    class InputNodesEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: int
        value: NodeData
        def __init__(self, key: _Optional[int] = ..., value: _Optional[_Union[NodeData, _Mapping]] = ...) -> None: ...
    INPUT_NODES_FIELD_NUMBER: _ClassVar[int]
    input_nodes: _containers.MessageMap[int, NodeData]
    def __init__(self, input_nodes: _Optional[_Mapping[int, NodeData]] = ...) -> None: ...

class NodeData(_message.Message):
    __slots__ = ("node_id", "function_param", "node_structure")
    NODE_ID_FIELD_NUMBER: _ClassVar[int]
    FUNCTION_PARAM_FIELD_NUMBER: _ClassVar[int]
    NODE_STRUCTURE_FIELD_NUMBER: _ClassVar[int]
    node_id: int
    function_param: FunctionParams
    node_structure: NodeConnector
    def __init__(self, node_id: _Optional[int] = ..., function_param: _Optional[_Union[FunctionParams, _Mapping]] = ..., node_structure: _Optional[_Union[NodeConnector, _Mapping]] = ...) -> None: ...

class FileTree(_message.Message):
    __slots__ = ("tree_type", "root_node", "is_default")
    TREE_TYPE_FIELD_NUMBER: _ClassVar[int]
    ROOT_NODE_FIELD_NUMBER: _ClassVar[int]
    IS_DEFAULT_FIELD_NUMBER: _ClassVar[int]
    tree_type: int
    root_node: NodeData
    is_default: bool
    def __init__(self, tree_type: _Optional[int] = ..., root_node: _Optional[_Union[NodeData, _Mapping]] = ..., is_default: bool = ...) -> None: ...

class apmFile(_message.Message):
    __slots__ = ("trees", "usr_id", "character_id")
    TREES_FIELD_NUMBER: _ClassVar[int]
    USR_ID_FIELD_NUMBER: _ClassVar[int]
    CHARACTER_ID_FIELD_NUMBER: _ClassVar[int]
    trees: _containers.RepeatedCompositeFieldContainer[FileTree]
    usr_id: int
    character_id: int
    def __init__(self, trees: _Optional[_Iterable[_Union[FileTree, _Mapping]]] = ..., usr_id: _Optional[int] = ..., character_id: _Optional[int] = ...) -> None: ...

class FunctionParams(_message.Message):
    __slots__ = ("function_prompt", "input_data_obj", "input_text", "system_prompt")
    FUNCTION_PROMPT_FIELD_NUMBER: _ClassVar[int]
    INPUT_DATA_OBJ_FIELD_NUMBER: _ClassVar[int]
    INPUT_TEXT_FIELD_NUMBER: _ClassVar[int]
    SYSTEM_PROMPT_FIELD_NUMBER: _ClassVar[int]
    function_prompt: str
    input_data_obj: bytes
    input_text: str
    system_prompt: str
    def __init__(self, function_prompt: _Optional[str] = ..., input_data_obj: _Optional[bytes] = ..., input_text: _Optional[str] = ..., system_prompt: _Optional[str] = ...) -> None: ...
