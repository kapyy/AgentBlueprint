import logging

from .function_reflect_gen import MainServicerDistributorCaller, SerializeDataWithTypeID, \
    SubServicerDistributionCaller
from .factory_util import isMinorFunction

def DeserializeFunctionNode(node):
    function_prompt = node.function_param.function_prompt
    system_prompt = node.function_param.system_prompt
    # immediate call if no child
    if isMinorFunction(node.node_id):
        if len(node.node_structure.input_nodes) == 0:
            logging.getLogger().error("Minor Function has no input nodes")
        # this index starts from 1 to n to match prompt format
        for index in node.node_structure.input_nodes:
            result, result_str = DeserializeFunctionNode(node.node_structure.input_nodes[index])
            function_prompt = function_prompt.replace("{" + str(index) + "}", " " + result_str)
            system_prompt = system_prompt.replace("{" + str(index) + "}", " " + result_str)
    context = {
        "prompt": function_prompt,
        "text_input": node.function_param.input_text,
        "data_input": node.function_param.input_data_obj, #dataInput is a formatted object that cannot be used as a string
        "system": system_prompt
    }
    result, result_str = MainServicerDistributorCaller(node.node_id, context)
    return result, result_str


def CallSubordinateFunction(function_id, data_id, data):
    serialized_data = SerializeDataWithTypeID(data, data_id)
    return SubServicerDistributionCaller(function_id, serialized_data)
