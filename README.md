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
- 
---

