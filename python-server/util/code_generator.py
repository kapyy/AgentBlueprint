import os.path
import re

import yaml
config_file_path = './conf.ini'
function_gen_path = 'factory/function_reflect_gen.py'
temp_resp_gen_path = 'prompt_template/response_format_gen.py'
func_implement_path = 'function/function_servicer_implementation.py'
static_ops_path = 'ops'
data_yaml_path = 'config/DataIndexGen.yaml'
func_yaml_path = 'config/FunctionIndexGen.yaml'

def _camel_to_snake(name):
    return re.sub('(.)([A-Z])', r'\1_\2', name).lower()
def _is_base_type(type):
    return type in ['int32', 'int64','uint32','uint64', 'float', 'double', 'string', 'bool']
def _is_number_type(type):
    return type in ['int32', 'int64','uint32','uint64', 'float', 'double']
def _is_string_type(type):
    return type in ['string']
def gen_function_reflect():

    with (open(os.path.join(function_gen_path), 'w+') as file):
        title = f'''
import configparser
import logging
import grpc
import message.data.dataIndexGen_pb2
from message.data.functionDistribute_pb2 import GeneralPyRequest
from message.data.functionDistribute_pb2_grpc import APMFunctionsServiceStub
config_parser = configparser.ConfigParser()
config_parser.read('{config_file_path}')

'''
        file.write(title)
        # Input Data
        file.write('def SerializeDataWithTypeID(data, data_id):\n')
        file.write('    if data_id == 0:\n')
        file.write('        return "No Data Type"\n')
        dataConf = yaml.load(open(data_yaml_path), Loader=yaml.FullLoader)
        dataList = dataConf['PluralDataIndex']
        dataList.update(dataConf['SingularDataIndex'])
        for (key,data) in dataList.items():
            lowercase_name = key.lower()
            file.write(f'    elif data_id == {data["index"]}:\n')
            file.write(f'        {lowercase_name} = message.data.dataIndexGen_pb2.{key}()\n')
            file.write(f'        {lowercase_name}.ParseFromString(data)\n')
            file.write(f'        return {lowercase_name}\n')
        file.write('    else:\n')
        file.write('        return None\n')
        file.write('\n')
        # DistributeCaller
        funcConf = yaml.load(open(func_yaml_path), Loader=yaml.FullLoader)
        file.write('def StaticFunctionDistributionCaller(id, rqst):\n')
        file.write('    with grpc.insecure_channel(config_parser[\'server.setting\'][\'FunctionURL\']) as channel:\n')
        file.write('        stub = APMFunctionsServiceStub(channel)\n')
        for (key,func) in funcConf['Functions'].items():
            if func['Type'] == 'StaticFunction':
                file.write(f'        if id == {func["FunctionID"]}:\n')
                file.write(f'            return stub.{key}(rqst)\n')
        file.write('\n')
        # Main function
        file.write('def MainServicerDistributorCaller(id, context):\n')
        file.write('    with grpc.insecure_channel(config_parser[\'server.setting\'][\'FunctionURL\']) as channel:\n')
    # try:
    #     grpc.channel_ready_future(channel).result(timeout=10)
    # except grpc.FutureTimeoutError:
    #     logging.error("Error connecting to py function server at %s", config_parser['server.setting']['FunctionURL'])
    #     return None, "No Data Type"
    # stub = APMFunctionsServiceStub(channel)
        file.write('        try:\n')
        file.write('            grpc.channel_ready_future(channel).result(timeout=10)\n')
        file.write('        except grpc.FutureTimeoutError:\n')
        file.write('            logging.error("Error connecting to py function server at %s", config_parser[\'server.setting\'][\'FunctionURL\'])\n')
        file.write('            return None, "No Data Type"\n')
        file.write('        stub = APMFunctionsServiceStub(channel)\n')
        for (key,func) in funcConf['Functions'].items():
            if func['Type'] == 'DefaultFunction':
                file.write(f'        if id == {func["FunctionID"]}:\n')
                file.write(f'            request = GeneralPyRequest(prompt=context["prompt"],text=context["text_input"],system_prompt=context["system"])\n')
                file.write(f'            return stub.{key}(request), "No Data String Out needed"\n')
        file.write('        else:\n')
        file.write('            logging.error("No such service Function found")\n')
        file.write('            return None, "No Data Type"\n')

def _return_data_map_with_name(name, data_map) -> str:
    map_desc = '{'
    for (data_key, data_val) in data_map.items():
        if data_key != name:
            continue
        for (field_key, field_val) in data_val['property'].items():
            field_type = field_val['type']
            if not _is_base_type(field_type):
                field_type = _return_data_map_with_name(field_type, data_map)
            map_desc += f'"{_camel_to_snake(field_key)}": {field_type},'
    map_desc += '}'
    return map_desc
def _return_data_map(index, data_map) -> str:
    map_desc = '{'
    for (data_key, data_val) in data_map.items():
        if data_val['index'] != index:
            continue
        for (field_key, field_val) in data_val['property'].items():
            field_type = field_val['type']
            if not _is_base_type(field_type):
                field_type = _return_data_map_with_name(field_type, data_map)
            map_desc += f'"{_camel_to_snake(field_key)}": {field_type},'
    map_desc += '}'
    return map_desc
def _return_prop_keys(index, data_map) -> str:
    list_desc = ''
    for (data_key, data_val) in data_map.items():
        if data_val['index'] != index:
            continue
        for (field_key, field_val) in data_val['property'].items():
            list_desc += f'"{_camel_to_snake(field_key)}",'
    return list_desc
def template_response_gen():
    with (open(os.path.join(temp_resp_gen_path), 'w+') as file):
        data_yaml = yaml.load(open(data_yaml_path), Loader=yaml.FullLoader)
        data_dict = data_yaml['PluralDataIndex']
        data_dict.update(data_yaml['SingularDataIndex'])
        data_dict.update(data_yaml['InternalDataIndex'])
        func_yaml = yaml.load(open(func_yaml_path), Loader=yaml.FullLoader)
        for (func_key,func_val) in func_yaml['Functions'].items():
            resp_data_index = func_val['OutputNode']
            file.write(f'{_camel_to_snake(func_key)}_response = """\n')
            file.write(f'Return a JSON list with objects containing {_return_prop_keys(resp_data_index,data_dict)} keys.\n')
            file.write('Example format are as follows:\n')
            file.write('[\n')
            for i in range(3):
                file.write(f'    {_return_data_map(resp_data_index, data_dict)}')
            file.write(']\n')
            file.write('"""\n')


def function_implement():
    with (open(os.path.join(func_implement_path), 'w+') as file):
        data_yaml = yaml.load(open(data_yaml_path), Loader=yaml.FullLoader)
        data_dict = data_yaml['PluralDataIndex']
        data_dict.update(data_yaml['SingularDataIndex'])
        data_dict.update(data_yaml['InternalDataIndex'])
        func_yaml = yaml.load(open(func_yaml_path), Loader=yaml.FullLoader)
        file.write('import logging\n')
        file.write('import message.data.dataIndexGen_pb2\n')
        file.write('import message.data.functionDistribute_pb2_grpc\n')
        file.write('from prompt_template import prompt_template, response_format_gen, system_template\n')
        file.write('from ops.jsonify_data import parse_json_list\n')
        file.write('from ops.pipe_util import get_llm_op, get_prompt\n')
        file.write('from util.service_config import ServiceConfig\n')
        file.write('\n')
        file.write('logger = logging.getLogger()\n')
        file.write('config = ServiceConfig()\n')
        file.write('config.embedding_device = \'cuda\'\n')
        file.write('\n')
        file.write('class APMFunctionsServiceServicer(object):\n')
        file.write('    """Internal Python Service to distribute the apm request to individual functions\n')
        file.write('    """\n')
        for (func_key,func_val) in func_yaml['Functions'].items():
            if func_val['Type'] == 'DefaultFunction':
                file.write(f'    def {func_key}(self, request, context):\n')
                file.write(f'        prompt_args = {{\n')
                file.write(f'            "system": system_template.gpt_system_prompt,\n')
                file.write(f'            "context": request.text,\n')
                file.write(f'            "prompt": request.prompt,\n')
                file.write(f'            "response": response_format_gen.{_camel_to_snake(func_key)}_response\n')
                file.write(f'        }}\n')
                file.write(f'        prompt_formatter = get_prompt(prompt_template.default_prompt)\n')
                file.write(f'        llm_caller = get_llm_op(config=ServiceConfig())\n')
                file.write(f'        # run pipeline\n')
                file.write(f'        formatted_prompt = prompt_formatter(prompt_args)\n')
                file.write(f'        result = llm_caller(\n')
                file.write(f'            formatted_prompt,\n')
                file.write(f'        )\n')
                file.write(f'        data = parse_json_list(result)\n')
                for (data_key,data_val) in data_dict.items():
                    if data_val['index'] == func_val['OutputNode']:
                        if func_val['Type'] == 'DefaultFunction':
                            file.write(f'        {data_key.lower()}_list = message.data.dataIndexGen_pb2.{data_key}List()\n')
                            file.write(f'        if data is None:\n')
                            file.write(f'            return {data_key.lower()}_list\n')
                            file.write(f'        for {data_key.lower()} in data:\n')
                            for (field_key,field_val) in data_val['property'].items():
                                type_example = None
                                if _is_number_type(field_val['type']):
                                    type_example = 0
                                elif _is_string_type(field_val['type']):
                                    type_example = '""'
                                file.write(f'            {_camel_to_snake(field_key)} = {data_key.lower()}.get("{_camel_to_snake(field_key)}", {type_example})\n')
                            file.write(f'            {data_key.lower()} = message.data.dataIndexGen_pb2.{data_key}(\n')
                            for (field_key,field_val) in data_val['property'].items():
                                file.write(f'                {_camel_to_snake(field_key)}={_camel_to_snake(field_key)},\n')
                            file.write(f'            )\n')
                            file.write(f'            {data_key.lower()}_list.{data_key.lower()}_list.append({data_key.lower()})\n')
                            file.write(f'        return {data_key.lower()}_list\n')
                file.write('\n\n')
            elif func_val['Type'] == 'StaticFunction':
                with (open(os.path.join(static_ops_path,_camel_to_snake(f'{func_key}Handler.py')),'w+') as handler_file):
                    handler_file.write('import message.data.dataIndexGen_pb2\n\n')
                    # def StringToEmoji(
                    #         content: message.data.dataIndexGen_pb2.Action) -> message.data.dataIndexGen_pb2.ParsedAction:
                    input_type,output_type = None,None
                    for (data_key,data_val) in data_dict.items():
                        if data_val['index'] == func_val['InputNode']:
                            input_type = data_key
                        if data_val['index'] == func_val['OutputNode']:
                            output_type = data_key
                    handler_file.write("#####################################\n"
                                       "#NEED IMPLEMENTATION FOR STATIC FUNCTION\n"
                                       "#####################################\n\n")
                    handler_file.write(f'def {func_key}Handler(request: message.data.dataIndexGen_pb2.{input_type}) -> message.data.dataIndexGen_pb2.{output_type}:\n')

                    handler_file.write(f'    return message.data.dataIndexGen_pb2.{output_type}()\n')
                file.write(f'    def {func_key}(self, request, context):\n')

                file.write(f'        from ops.{_camel_to_snake(f"{func_key}Handler")} import {func_key}Handler\n')
                file.write(f'        response = {func_key}Handler(request)\n')
                file.write(f'        return response\n')
