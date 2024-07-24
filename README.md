# AgentBlueprint
***
## What is AgentBlueprint
AgentBlueprint is a base framework that deeply connects between application code and LLM.
## What does AgentBlueprint have
AgentBlueprint has roughly 3 parts:
1. **Single Agent Blueprint**: A blueprint is a pre-defined structure of Data/Prompts that represents how the agent should behave in a controllable fashion.
2. **Data-Driven Structure**: The whole framework is data-driven, meaning that most of the data variable that might be used later in the LLM is pre-generated in the framework. This allows a frictionless connection between the application data and the LLM.
3. **Multi-Agent Support**: Because of the data-driven structure and self-determined nature of each agent, multi-agent support is a natural feature of the framework. 

## Current State
- [X] Game Agent
- [ ] Python Support
- [ ] General Agent Functionalities (Copilot, etc.)
- [ ] Blueprint Node-Graph Edtor

---
# Getting Started
1. There are two repositories in this project, one is golang and one is python. Installing both environments are required to run the project (for now).
2. install protoc with grpc plugin for python and golang
```python
pip install grpcio-tools
```
```golang
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
3. The whole project is build upon data-driven structure, so the first thing to do is to define all the data/ function(data pipline) under `/DataConfig[EDIT ME]/`  that might be used later. See Section [] for detail explanation.
4. Run Generate.bat for windows.
5. There are few codes that need to be implemented.
- Implement the data source at `golang-client/implementation/impl_gen_XXX`
  >when the data structure is generate, the data source needs to be connected to the original source, could be a database,RAG,runtime variable,temporary cache, etc.

6. Write APM (Agent Personality Module).
> APM file is the core blueprint of each agent, it contains every task's prompt. APM assembles runtime data with a data placeholder and prompt to execute a complete query.
7. Initialize agent at `golang-client/client_object/agent_entity.go` and deserialize the pre-defined APM to load to the AgentEntity by calling `DeserializeAPMToEntity` function.
8. Run the generated pipelines at `golang-client/client_object/entity_function_gen.go` returns the result of the Query.
---
## Step by Step Guide On how to use the framework
#### To Build a Generative Agent Multi Agent Demo
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



---
## FAQ
1. **What does DataInstance Do?** <br>
Data Instance is the data structure of everything, it contains the data its self, context information, and chaining information of where this data is from (always another DataInstance). So you can think it as the equivalent replacement to a data variable within the agent class.
> In our vision, this Project will be built upon Data Instance structure entirely. DataInstance is the intermediate structure that connects between blueprint's runtime data call, LLM's query result, database's data retrieval. It more of a ECS(EntityComponentSystem) design pattern.
2. **Why uses Data-Oriented Design?** <br>
There are many reasons for that:
   - *More controllable and retrievable data.*<br> In an AI app development, the result that LLM returned are often used later in the application, so having a data structure that can be easily retrieved and manipulated is very handy. 
   - *RAG support*<br> Though this project has not shown any rag support yet. I have try to play around with RAG, and it seems that a clear data structure boosts retrieval's quality a lot from a graph point of view. Also, when it comes to graph Indexing, if a structure can be provided beforehand, the indexing process will adapt to application's data structure frictionlessly.
   - *Context Awareness*<br> A unique context reference is connected to each data instance. When it comes to retrieving logical connections between data, the context will work as a tree that can be easily traversed. The Tree of Thoughts/ Chain of Thoughts is the most common example of this. Context design also frees up the task from being strictly sequential; it can be done in parallel or execut from desire node in the past.
   - *Blueprint Support*<br> This is where this project firstly started, we were trying to do a runtime data query with LLM at first. but to retrieve different data dynamically requires a tremendous amount of work. So we decided to build common data interface that adapts to the data indexing with whom the assembly can happen at resource level during runtime.
3. **What Is .apm File** <br>
APM or .apm file is a unique file format we defined, when I refer to Agent Blueprint / Assembly Structure. I mean this. It basically holds prompt information, with protobuf encoding. Each tasks can hold one or more unique LLM/LM prompts, and in that prompt all the runtime/static data is replaced by a placeholder. where assembly happens when the task is called, then the placeholder will be replaced by the actual data. It's something that looks like this:
    ```
    "{agent_information} \n" + \
    "Yesterday activities: {daily_summary} \n" + \
    "In addition, he has a special mission today, {today's mission}" + \
    "with the above information provide a plan that this agent will do \n" + \
    ```
   or 
    ```
    Currently {agent_name} saw {interact_object_name}is in {interact_object_state} State,
    Now {agent_name} is at {agent_location} in the {current_time},
    What would he do next to react to this current situation?
   ```
    Even with the same task, to provide different .apm assemblies(blueprint) will result in different queries. This is the core of the blueprint structure - to make every task exchangeable and customizable.
4. 