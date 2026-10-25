"""Minimal chat completion against APIClaw with the official OpenAI Python SDK.

pip install openai
export APICLAW_API_KEY=...  APICLAW_MODEL=...
"""
import os

from openai import OpenAI

client = OpenAI(
    base_url="https://apiclaw.biz/v1",
    api_key=os.environ["APICLAW_API_KEY"],
)

resp = client.chat.completions.create(
    model=os.environ.get("APICLAW_MODEL", "YOUR_MODEL_ID"),
    messages=[{"role": "user", "content": "Say hello in one sentence."}],
)
print(resp.choices[0].message.content)

# Streaming
stream = client.chat.completions.create(
    model=os.environ.get("APICLAW_MODEL", "YOUR_MODEL_ID"),
    messages=[{"role": "user", "content": "Count to five."}],
    stream=True,
)
for chunk in stream:
    if chunk.choices and chunk.choices[0].delta.content:
        print(chunk.choices[0].delta.content, end="", flush=True)
print()
