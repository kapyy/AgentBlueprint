
import message.data.dataIndexGen_pb2
import message.data.functionDistribute_pb2_grpc
from ops.string_to_emoji import StringToEmoji
from prompt_template import template, response_format, system

from ops.jsonify_data import parse_json_list


import logging

from ops.pipe_util import get_llm_op, get_prompt
from util.service_config import SerivceConfig

logger = logging.getLogger("kotoko.debug")

config = SerivceConfig()
config.embedding_device = 'cuda'

class APMFunctionsServiceServicer(object):
    """Internal Python Service to distribute the apm request to individual functions
    """

    def InsertActionsWithObservation(self, request, context):
        prompt_args = {
            "system": system.gpt_system_prompt_without_background,
            "context": request.text,
            "prompt": request.prompt,
            "response": response_format.insert_actions_with_observation_response
        }
        prompt_formatter = get_prompt(template.insert_actions_with_observation_prompt)
        llm_caller = get_llm_op(config=SerivceConfig())
        # run pipeline
        formatted_prompt = prompt_formatter(prompt_args)
        result = llm_caller(
            formatted_prompt,
        )

        data = parse_json_list(result)
        actions = message.data.dataIndexGen_pb2.ActionList()
        if data is None:
            return actions
        for action in data:
            action_desc = action.get("action_description", "")
            duration = action.get("duration", 0)
            action = message.data.dataIndexGen_pb2.Action(
                action_description=action_desc,
                duration=duration,
            )
            actions.action_list.append(action)
        return actions


    def ActionFormatter(self, request, context):
        fmt_action = message.data.dataIndexGen_pb2.ParsedAction()
        emoji_list = StringToEmoji(request.action_description)
        fmt_action.emoji_list = emoji_list
        return fmt_action
