# Generated Script
# Function/Data ID to proto/grpc Refection Dictionary

MESSAGE_REFLECTION = {

    1001: "BaseObject",
    1002: "MemoryLog",
    1003: "Plan",
    1004: "ActionData",
    1005: "Event",
    1006: "ActionLog",
    1015: "ChatLog",
    1016: "TalkLog",
    2001: "BaseObjects",
    2003: "Plans",
    2004: "ActionDatas",
    3001: "FormattedObject",
    3002: "EmojiData",
    3003: "LinearIndicator",
    3004: "Skill",
    3005: "StoryLog",
    3006: "ObservationNeedToReact",
    4001: "CharacterBaseInfo",
    4002: "WhisperContent",
    4003: "CoreValues",
    4004: "CorePersonality",
    4005: "SkillSet",
    4006: "StoryModuleLogs",
    4007: "DispatchModuleLog",
    4008: "Item",
    4009: "ItemStatus",
    4010: "Items",
    4011: "ObservedSituations",
    4012: "ObservedSituation",
    4013: "ScoredMemory",
    4014: "FormattedAction",
    4015: "RelevantInfo",
    4016: "ParsedAction",
    4017: "GeneralString",
    4028: "HungerIntent",
    4029: "FatigueIntent",
}


def getMessageName(id):
    return "apm." + MESSAGE_REFLECTION[id]

