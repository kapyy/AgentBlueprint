1. **Define Data Structure** <br>
   Generative agent have plans, execute its plans, reflect it to memory, watch other people action and generate its own plans<br>
   so we need to define the data structure for each of them.
      - AgentBackground
        - Name
        - Background
      - Plan
        - description
      - ExecuteResult
        - isSucceed
        - description
      - Memory
        - description 
        
2. **Define Function Tasks**<br>
then we defined a few tasks that this agent will have
      - generate plans 
        - generate plans from nothing
        - generate plans from other's plans
      - execute plans
      - reflect to memories
3. write to two .yaml
    >The difference between PluralDataIndex and SingularDataIndex are that data belongs to plural when it needs to be returned as a list from a task, like if a task generate a list of actions, then action is a pluralDataIndex   
    
    >The difference between Default Function and Static Function are that static functions will only take very specific input, and its functionality is fixed and implemented at python hardcoded. Default function should be used most of the time. and its input will be implemented in the apm file.
