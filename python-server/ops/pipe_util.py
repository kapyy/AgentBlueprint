import logging

# from groq import Groq
from typing import Dict
from openai import OpenAI
from typing import List
# from monitors.api import log_openai_token

# def _get_prompt(template):
#     return MainServicePrompt(prompt_temp=template)

def _parse_input_openai(messages: List[dict]):
    assert isinstance(messages, list)
    new_messages = []
    for m in messages:
        if not ('role' in m and 'content' in m):
            raise KeyError("Invalid message keys: 'role' and 'content' must be in message")

        if m['role'] in ['system', 'assistant', 'user']:
            new_messages.append(m)
        else:
            if m['role'] == 'question':
                new_m = {'role': 'user', 'content': m['content']}
            elif m['role'] == 'answer':
                new_m = {'role': 'assistant', 'content': m['content']}
            else:
                raise KeyError('Invalid message role: only accept role value from ["question", "answer", "system"].')
            new_messages.append(new_m)
    return new_messages


# LLM operation singleton
class LLMOp:
    def __init__(self):
        self.llm_op_groq = None
        self.llm_op_openai = None
        self.llm_op_dolly = None

    def get_llm_op(self, config):
        if config.customize_llm:
            return config.customize_llm

        # if config.llm_src.lower() == "groq":
        #     if self.llm_op_groq:
        #         return self.llm_op_groq
        #     else:
        #         def altfunc(f, source_pipeline: str = "Unknown pipeline"):
        #             logger = logging.getLogger("kotoko.llm")
        #             logger.debug("[llm ->] %s invoked with groq", str(f))
        #             client = Groq(api_key = config.groq_api_key)
        #             x = client.chat.completions.create(
        #                 {"role": "system",
        #                 "content": "Role Play: All your actions should consider to the personality given in the prompt"},
        #                 messages=[
        #                     {"role": i, "content": f[0][i]} for i in f[0]
        #                 ],
        #                 model=config.groq_model
        #             ).choices[0].message.content
        #             logger.debug("[llm <-] %s", str(x))
        #             return x
        #         self.llm_op_groq = altfunc
        #         return altfunc
        #
        # el
        if config.llm_src.lower() == "openai":
            if self.llm_op_openai:
                return self.llm_op_openai
            else:
                # f: List[Dict[str, str]] messages for openai completion agent
                # source_pipeline: str source pipeline
                def altfunc(f, source_pipeline: str = "Unknown pipeline"):
                    logger = logging.getLogger("kotoko.llm")
                    logger.debug("[llm ->] %s invoked with openai", str(f))
                    # prepare client prompts and contexts
                    client = OpenAI(api_key=config.openai_api_key)
                    unparsed_messages = [{"role": i, "content": f[0][i]} for i in f[0]]
                    # unparsed_messages.append({"role": "system", "content": "Role Play: All your actions should consider to the personality given in the prompt"})
                    client_messages = _parse_input_openai(messages=unparsed_messages)

                    # call openai api
                    api_response = client.chat.completions.create(
                        messages=client_messages,
                        model=config.openai_model
                    )

                    # monitor token usage
                    prompt_tokens_used = api_response.usage.prompt_tokens
                    completion_tokens_used = api_response.usage.completion_tokens
                    total_tokens_used = api_response.usage.total_tokens
                    # logger.debug("[TOKEN usage]\nsource_pipeline: %s\nprompt_tokens: %s\ncompletion_tokens: %s\ntotal_tokens: %s\n",
                    #              source_pipeline, prompt_tokens_used, completion_tokens_used, total_tokens_used)
                    #log_openai_token(source_pipeline, prompt_tokens_used, completion_tokens_used, total_tokens_used)

                    # return the completion
                    completion = api_response.choices[0].message.content
                    logger.debug("[llm <-] %s", str(completion))
                    return completion
                return altfunc


        raise RuntimeError('Unknown llm source: [%s], only support groq, openai and dolly' % (config.llm_src))


llm_ops = LLMOp()

gpt_prompt = """
{context}
"""
class ServicePrompt:
    def __init__(self, prompt_temp:str = None):
        if prompt_temp:
            self.prompt_temp = prompt_temp
        else:
            self._template = gpt_prompt

    def __call__(self, context: Dict[str, str]):
        prompt_str = self._template.format(context=context['context'], prompt=context['prompt'])
        ret = [{'system': context['system'], 'question': prompt_str, "answer": context['response']}]
        return ret
def get_llm_op(config):
    return llm_ops.get_llm_op(config)

def get_prompt(template):
    return ServicePrompt(prompt_temp=template)
