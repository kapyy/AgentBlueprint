from google.protobuf.descriptor_pool import DescriptorPool
import grpc
from google.protobuf.message_factory import GetMessageClass
from grpc_reflection.v1alpha.proto_reflection_descriptor_database import (
    ProtoReflectionDescriptorDatabase,
)

from .dict_data_gen import getMessageName


################ NOT IN USE ################
# reflection are done directly from switch case
#
# def run():
#     print("Will try to greet world ...")
#     with grpc.insecure_channel("localhost:50077") as channel:
#         reflection_db = ProtoReflectionDescriptorDatabase(channel)
#         services = reflection_db.get_services()
#         print(f"found services: {services}")
#
#         desc_pool = DescriptorPool(reflection_db)
#         service_desc = desc_pool.FindServiceByName(getRPCService())
#         print(f"found Greeter service with name: {service_desc.full_name}")
#         for methods in service_desc.methods:
#             print(f"found method name: {methods.full_name}")
#             input_type = methods.input_type
#             print(f"input type for this method: {input_type.full_name}")
#             print()
#
#         request_desc = desc_pool.FindMessageTypeByName(
#             getMessageName(2003)
#         )
#         print(f"found request name: {request_desc.full_name}")

def _getMessagePrototypeFromID(id):
    with grpc.insecure_channel("localhost:50077") as channel:
        reflection_db = ProtoReflectionDescriptorDatabase(channel)
        desc_pool = DescriptorPool(reflection_db)
        message_desc = desc_pool.FindMessageTypeByName(
            getMessageName(id)
        )
        return GetMessageClass(message_desc)


def _getMessagePrototypeFromName(name):
    with grpc.insecure_channel("localhost:50077") as channel:
        reflection_db = ProtoReflectionDescriptorDatabase(channel)
        desc_pool = DescriptorPool(reflection_db)
        message_desc = desc_pool.FindMessageTypeByName(
            name
        )
        return GetMessageClass(message_desc)


def deserializeData(id, raw_data):
    data_obj = _getMessagePrototypeFromID(id)()
    data_obj.ParseFromString(raw_data)
    return data_obj


def deserializeDataWithName(name, raw_data):
    data_obj = _getMessagePrototypeFromName(name)()
    data_obj.ParseFromString(raw_data)

    return data_obj


