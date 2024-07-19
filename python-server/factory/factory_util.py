def isMinorFunction(id):
    # DefaultFunction --1
    # MinorFunction --2
    # StaticFunction --3 with static input and output
    return id // 10**8 == 2