# Vercel AI SDK + APIClaw

```bash
npm install ai @ai-sdk/openai-compatible
APICLAW_API_KEY=YOUR_API_KEY APICLAW_MODEL=apiclaw/claude-sonnet-5.5 npx tsx index.ts
```

Alternatively, with `@ai-sdk/openai`: `createOpenAI({ baseURL: "https://apiclaw.biz/v1", apiKey })` and call `apiclaw.chat("apiclaw/claude-sonnet-5.5")` to force the Chat Completions endpoint.
