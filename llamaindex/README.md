# LlamaIndex + APIClaw

```bash
pip install llama-index-llms-openai-like
export APICLAW_API_KEY=YOUR_API_KEY APICLAW_MODEL=apiclaw/claude-sonnet-5.5
python main.py
```

Use `OpenAILike` (not `OpenAI`) so LlamaIndex doesn't validate the model name against OpenAI's list. Set `context_window` to your model's real value from the [model hub](https://apiclaw.biz/ui/model_hub/).
