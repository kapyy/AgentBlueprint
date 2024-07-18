import json
import logging
import re


logger = logging.getLogger("kotoko.debug")

def parse_json(res):
    tmp_json_str = ""
    res_json_data = {}
    for item in res:
        if item == "{":
            tmp_json_str = item
            continue
        tmp_json_str += item
        if item == "}":
            try:
                res_json_data = json.loads(tmp_json_str)
            except json.JSONDecodeError:
                logger.debug("JSONDecodeError: Failed to parse JSON data. Source data: %s", tmp_json_str)
                res_json_data = None
            break
        
    return res_json_data

def parse_json_list(res):
    tmp_json_str = ""
    json_list = []
    for item in res:
        if item == "{":
            tmp_json_str = item
            continue
        tmp_json_str += item
        if item == "}":
            try:
                res_json_data = json.loads(tmp_json_str)
                json_list.append(res_json_data)
            except json.JSONDecodeError:
                logger.debug("JSONDecodeError: Failed to parse JSON data. Source data: %s", tmp_json_str)
                pass  # skip if failed
            tmp_json_str = ""

    return json_list

def extract_numbers(text):
    if isinstance(text, str):
        numbers = re.findall(r'\d+', text)
        return [int(num) for num in numbers]
    elif isinstance(text, int):
        return [text]