"""LlamaIndex OpenAILike pointed at APIClaw.

pip install llama-index-llms-openai-like
"""
import os

from llama_index.llms.openai_like import OpenAILike

llm = OpenAILike(
    model=os.environ.get("APICLAW_MODEL", "apiclaw/claude-sonnet-5.5"),
    api_base="https://apiclaw.biz/v1",
    api_key=os.environ["APICLAW_API_KEY"],
    is_chat_model=True,
    is_function_calling_model=True,
    context_window=128000,  # set to your model's real context window
)

print(llm.complete("Explain vector search in one sentence."))
