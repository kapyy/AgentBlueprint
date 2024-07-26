import logging
import json
import pandas as pd
import message.proto.dataIndexGen_pb2
import message.proto.functionDistribute_pb2_grpc
from prompt_template import prompt_template, response_format_gen, system_template
from ops.jsonify_data import parse_json_list
from ops.pipe_util import get_llm_op, get_prompt
from util.service_config import ServiceConfig

logger = logging.getLogger()
config = ServiceConfig()
config.embedding_device = 'cuda'


class APMFunctionsServiceServicer(object):
    """Internal Python Service to distribute the apm request to individual functions
    """

    def InsertActionWithObservation(self, request, context):
        prompt_args = {
            "system": system_template.gpt_system_prompt,
            "context": request.text,
            "prompt": request.prompt,
            "response": response_format_gen.insert_action_with_observation_response
        }
        prompt_formatter = get_prompt(prompt_template.default_prompt)
        llm_caller = get_llm_op(config=ServiceConfig())
        # run pipeline
        formatted_prompt = prompt_formatter(prompt_args)
        result = llm_caller(
            formatted_prompt,
        )
        try:
            data = json.loads(result)
        except json.JSONDecodeError:
            logger.error("JSONDecodeError: Failed to parse JSON data. Source data: %s", result)
            data = None
        if data is None:
            return message.proto.dataIndexGen_pb2.ParsedAction()
        df = pd.json_normalize(data)
        parsedaction = message.proto.dataIndexGen_pb2.ParsedAction(
            emoji_list=message.proto.dataIndexGen_pb2.EmojiData(emoji_description=df["emoji_list.emoji_description"],
                                                                emoji_unicode=df["emoji_list.emoji_unicode"], ), )
        return parsedaction

    def ActionFormatter(self, request, context):
        from ops.action_formatter_handler import ActionFormatterHandler
        response = ActionFormatterHandler(request)
        return response
