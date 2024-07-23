insert_action_with_observation_response = """
Return a JSON list with objects containing "action_description","duration","start_time","end_time", keys.
Example format are as follows:
[
    {"action_description": string,"duration": int32,"start_time": uint64,"end_time": uint64,}    {"action_description": string,"duration": int32,"start_time": uint64,"end_time": uint64,}    {"action_description": string,"duration": int32,"start_time": uint64,"end_time": uint64,}]
"""
action_formatter_response = """
Return a JSON list with objects containing "emoji_list", keys.
Example format are as follows:
[
    {"emoji_list": {"emoji_description": string,"emoji_unicode": string,},}    {"emoji_list": {"emoji_description": string,"emoji_unicode": string,},}    {"emoji_list": {"emoji_description": string,"emoji_unicode": string,},}]
"""
