# TODO: most API endpoints support constrained JSON generation. Fix this.

get_day_plan_prompt = """{context}
{prompt}

Given the Agent's information above, construct a daily routine for this agent starting from 00:00. The routine should span at least 24 hours.

Each entry in the schedule should include a timestamp and an action. For example: '1. 08:00-09:00 = Having breakfast at home'
"""

# fix this maybe
get_todo_list_prompt = """{context}
{prompt}

Given the Agent's information above, construct a todo list.
"""
get_inbetween_plan_prompt = """{context}
{prompt}

Provide a detailed schedule. Each plan should last at least 10 minutes. The format should be 'HH:MM-HH:MM = Activity', for example, '22:00-22:10 = Deep sleep at home'.
"""

# GPT might get confused about 0
execute_action_with_plans_prompt = """{context}
{prompt}

Based on the information above, what will the agent do next? The response should be in JSON and include the 'action_state' key, equal to 0 if the agent follows the provided plan, and otherwise 1.
"""

execute_action_with_whispers_prompt = """{context}
{prompt}

Based on the information above, what will the agent do next? If the new plan is reasonable, the agent prefers to follow it.
The response should be in JSON and include the 'action_state' key, equal to 0 if the agent follows the new plan, and otherwise 1.
"""
summarize_agent_prompt = """{context}
What you have done today is as follows:
{prompt}
Considering the given information, generate (a)a diary of agent's events within 140 characters; (b) three most important, memorable or unusual things that the agent has done
"""

observation_into_memory_log_prompt = """{context}
{prompt}

Generate a memory log about the observed situation.
"""

execute_action_with_observation_prompt = """{context}
{prompt}

Based on the above information, what action will the agent take next? 
"""

#generate_events_prompt = """{context}
# {prompt}
# Based on the information above, please generate a chronological list of events (three to five would be best) that the character needs to accomplish on a new day, each taking no less than 1h.
# The answer must be in this format: [prepare for the exam] [Conducting research and writing a detailed report for work]...[event_description] 
# Please describe what event the role will take in the "event_description" attribute, using the verb phrase format as the response.
# """

generate_events_prompt = """{context}
{prompt}
"""

get_chat_response_prompt = """{prompt}

Now, you just received the message '{context}'. 
Your reply can not contain emoji and don't use more than 50 words.
"""

generate_actions_from_event_prompt = """{context}
{prompt}
"""

insert_actions_with_whisper_prompt = """{prompt}
User Command: {context}
If you can obey the command,keep this command unchanged and return word by word, And the 'reason' key should be "" in string format.
If not, imagine an action you would do similar to this command instead, and give me a reason for the change.
The action description contains only one verb (for example, 'Have breakfast') and taking no more than 10 minutes.
Also generate a emotion that you would react to this command, from value 0 being upset/unhappy,1 being neutral and 2 being happy. return with key "emotion" and number value
"""
#
#

insert_actions_with_observation_prompt ="""
{prompt}

What you observed is: {context}\n
The action description only contains a verb (for example, 'Have breakfast'). 
"""

generate_hunger_habits_prompt = """{context}
{prompt}
"""
generate_fatigue_habits_prompt = """{context}
{prompt}
"""
#get_talk_content_prompt = """An agent is going to {context}.
# {prompt}
# Based on the information above, please generate conversations that may occur between the agent and the target. 
# Output the response in a JSON list, where each JSON object has two keys: 'id' and 'content'. Set 'id' to 0 for the agent's spoken words, and 1 for the target's spoken words. Only include the actual spoken content in the 'content' value, without any additional context or descriptors."
# """

# TODO: I SURE HOPE THIS IS OBSOLETED!!!

get_talk_content_prompt = """

{prompt}

Now, your talking topic is {context}.
Generate one or multiple rounds of conversation between the subject and the object.
"""
#Consider the possibility that the state of the related items may or may not change during the action. If there is no obvious change of the relevant item or if there is no relevant item, return "Idle"; else, using an adjective to describe the new obvious state of the mentioned relevant item.
#The response should be in JSON format and include the 'state' key.

action_executor_prompt="""{context}
{prompt}

If this action will be executed, how the visible state of related items will change? The response should be in JSON format and include the 'state' key.
If the action can change the relevant item's state, using an adjective to describe the new obvious state in the 'state' key; otherwise, it should be "" in string format".
"""

get_relevant_info_prompt="""{context}
{prompt}

Generate a brief summary based on the information above, especially the relations between nodes. 
"""

memory_distill_prompt="""{context}
{prompt}

Generate a brief summary based on the information above, especially the relations between nodes.
"""
dobit_summary_letter="""{context}
{prompt}
Write a letter to the [Player], summarizing today's activities, achievements, and potential emotional experiences. 
Frame the letter within the context of an intimate friendship. Ensure the letter is concise, with a maximum word count of 150 words and less than 200 tokens.

"""
dobit_recall_letter="""{context}
{prompt}
Write a letter to the [Player], focusing on expressing your feeling and thoughts about the day:
What you did at home today. These activities should be indoor activities.
Indoor life is boring, and you long to go outside and interact with other people.
Ensure the letter is concise, with a maximum word count of 150 words and less than 150 tokens.
"""
