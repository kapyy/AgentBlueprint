import logging
import message.data.dataIndexGen_pb2
import message.data.functionDistribute_pb2_grpc
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
        data = parse_json_list(result)
        action_list = message.data.dataIndexGen_pb2.ActionList()
        if data is None:
            return action_list
        for action in data:
            action_description = action.get("action_description", "")
            duration = action.get("duration", 0)
            start_time = action.get("start_time", 0)
            end_time = action.get("end_time", 0)
            action = message.data.dataIndexGen_pb2.Action(
                action_description=action_description,
                duration=duration,
                start_time=start_time,
                end_time=end_time,
            )
            action_list.action_list.append(action)
        return action_list


    def ActionFormatter(self, request, context):
        from ops.action_formatter_handler import ActionFormatterHandler
        response = ActionFormatterHandler(request)
        return response
