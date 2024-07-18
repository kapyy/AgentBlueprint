from apmFunctionFactory.apm_factory_dict_function_gen import MainServicerDistributorCaller, SerializeDataWithTypeID, \
    SubServicerDistributionCaller


def DeserializeFunctionNode(node):
    function_prompt = node.function_param.function_prompt
    system_prompt = node.function_param.system_prompt
    # immediate call if no child
    if node.HasField("node_structure"):
        # this index starts from 1 to n to match prompt format
        for index in node.node_structure.input_nodes:
            result, result_str = DeserializeFunctionNode(node.node_structure.input_nodes[index])
            # print("result: ", result, " index: ", index, " result_str: ", result_str)
            function_prompt = function_prompt.replace("{" + str(index) + "}", " " + result_str)
            system_prompt = system_prompt.replace("{" + str(index) + "}", " " + result_str)
            # print("prompt: ", prompt)
    node.function_param.function_prompt = function_prompt
    node.function_param.system_prompt = system_prompt
    context = {
        "prompt": node.function_param.function_prompt,
        "text_input": node.function_param.input_text,
        "data_input": node.function_param.input_data_obj,
        "system": node.function_param.system_prompt
    }
    result, result_str = MainServicerDistributorCaller(node.node_id, context)
    return result, result_str


def CallSubordinateFunction(function_id, data_id, data):
    serialized_data = SerializeDataWithTypeID(data, data_id)
    return SubServicerDistributionCaller(function_id, serialized_data)
