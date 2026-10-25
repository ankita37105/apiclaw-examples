# Aider + APIClaw

Aider supports any OpenAI-compatible API via `OPENAI_API_BASE`. Prefix the model name with `openai/`:

```bash
export OPENAI_API_BASE=https://apiclaw.biz/v1
export OPENAI_API_KEY=YOUR_API_KEY
aider --model openai/apiclaw/claude-sonnet-5.5
```

Or in `.aider.conf.yml`:

```yaml
openai-api-base: https://apiclaw.biz/v1
openai-api-key: YOUR_API_KEY
model: openai/apiclaw/claude-sonnet-5.5
```

Aider may warn about unknown context-window sizes for custom models. That's harmless, or you can define them in `.aider.model.metadata.json`.
