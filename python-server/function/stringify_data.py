#####################################################
######### DO NOT REMOVE THIS FILE ###################
# It is used when minor function are resolved within
# in function pipelines and need immediate return.
#####################################################

# def stringify_ScoredMemory(self):
#     print("stringify_ScoredMemory: ", self)
#     logs = ""
#     for i, mem in enumerate(self.scored_memories):
#         logs += ("Memory:" +
#                  f"start time:{str(mem.memory_log.start_time)} ," +
#                  f"end time:{str(mem.memory_log.end_time)} ," +
#                  f"action description:{mem.memory_log.action_description} ," +
#                  f"with a importance score of:{str(mem.score)} .")
#     print("ScoredMemories: ", logs)
#     return "MemoryLogs: {}".format(logs)
#
#
# def stringify_ObservationNeedToReact(self):
#     if self.observation_description == "":
#         return ""
#     logs = "What the character observes is following: " + self.observation_description + ". "
#     logs += "In response to what has been observed above, the character should take some actions!"
#     print("ObservationReaction: ", logs)
#     return logs
