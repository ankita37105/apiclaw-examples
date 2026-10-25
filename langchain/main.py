"""LangChain ChatOpenAI pointed at APIClaw.

pip install langchain-openai
"""
import os

from langchain_openai import ChatOpenAI

llm = ChatOpenAI(
    model=os.environ.get("APICLAW_MODEL", "apiclaw/claude-sonnet-5.5"),
    base_url="https://apiclaw.biz/v1",
    api_key=os.environ["APICLAW_API_KEY"],
)

print(llm.invoke("Give me one tip for writing clean Python.").content)
