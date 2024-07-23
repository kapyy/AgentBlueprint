import configparser
from typing import Optional, Any, Union

from pydantic import BaseModel

import os

class ServiceConfig(BaseModel):
    """
    Config of pipeline
    """

    # config for sentence_embedding
    embedding_model: Optional[str] = 'all-MiniLM-L6-v2'
    openai_api_key: Optional[str] = os.environ.get("OPENAI_API_KEY")
    groq_api_key: Optional[str] = os.environ.get("GROQ_API_KEY")
    groq_model: Optional[str] = "mixtral-8x7b-32768"

    embedding_device: Optional[int] = -1
    # config for search_milvus
    host: Optional[str] = '127.0.0.1'
    port: Optional[str] = '19530'
    collection_name: Optional[str] = 'chatbot'
    top_k: Optional[int] = 5
    user: Optional[str] = None
    password: Optional[str] = None
    # config for llm
    llm_src: Optional[str] = 'openai'
    openai_model: Optional[str] = 'gpt-3.5-turbo'
    dolly_model: Optional[str] = 'databricks/dolly-v2-3b'

    ernie_api_key: Optional[str] = None
    ernie_secret_key: Optional[str] = None

    customize_llm: Optional[Any] = None
    customize_prompt: Optional[Any] = None
    # config for rerank
    rerank: Optional[bool] = False
    rerank_model: Optional[str] = 'cross-encoder/ms-marco-MiniLM-L-6-v2'
    threshold: Optional[Union[float, int]] = 0.6
