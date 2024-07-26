pandastest_str = """
[
    {"action_description": "string","duration": "int32","start_time": "uint64","end_time":" uint64","emoji": {"emoji_description": "string","emoji_unicode": "string"}},
    {"action_description": "string","duration": "int32","start_time": "uint64","end_time":" uint64","emoji": {"emoji_description": "string","emoji_unicode": "string"}},
    {"action_description": "string","duration": "int32","start_time": "uint64","end_time":" uint64","emoji": {"emoji_description": "string","emoji_unicode": "string"}}
    ]
"""
import pandas as pd
import json

try :
    file  = json.loads(pandastest_str)
except ValueError as e:
    print(f"Error parsing JSON data: {e}")
    file = None
df = pd.json_normalize(file)
for index, row in df.iterrows():
    # print(row)
    print(row["action_description"])
    print(row["emoji.emoji_description"])
    print(row["emoji.emoji_unicode"])
    print(row["duration"])
    print(row["start_time"])
    print(row["end_time"])
    print("")