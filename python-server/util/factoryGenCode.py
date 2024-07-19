import yaml
config_file_path = './conf.ini'
def gen_function_reflect():

    with open('../factory/function_reflect_gen1.py', 'w') as file:
        title = f'''
            import configparser\n
            import logging\n
            import grpc\n
            import message.data.dataIndexGen_pb2\n
            from message.data.functionDistribute_pb2 import GeneralPyRequest\n
            from message.data.functionDistribute_pb2_grpc import APMFunctionsServiceStub\n
            config_parser = configparser.ConfigParser()\n
            config_parser.read({config_file_path})\n
        '''
        file.write(title)
        # Input Data
        file.write('def SerializeDataWithTypeID(data, data_id):\n')
        file.write('    if data_id == 0:\n')
        file.write('        return "No Data Type"\n')
        for

