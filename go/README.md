# Go + APIClaw (OpenAI-compatible, flat-rate)

A tiny, dependency-free Go client and CLI for **[APIClaw](https://apiclaw.biz)**, a flat-rate, OpenAI-compatible AI API gateway. One key and one base URL (`https://apiclaw.biz/v1`) reach Claude, GPT, Gemini, DeepSeek, Qwen, Kimi, GLM and more. Plans are $19–$129/month, billed by daily requests rather than tokens, and new accounts get 50 free trial requests ([sign up](https://apiclaw.biz/ui/signup/)).

```bash
go get github.com/ankita37105/apiclaw-examples/go
```

```go
c := apiclaw.NewClient("") // reads APICLAW_API_KEY
reply, err := c.Ask(ctx, "apiclaw/claude-sonnet-5.5", "Say hello in five words.")
```

CLI:

```bash
export APICLAW_API_KEY=YOUR_API_KEY
go run github.com/ankita37105/apiclaw-examples/go/cmd/apiclaw-chat@latest "Explain HTTP 429 in one line"
```

- Default model: `apiclaw/claude-sonnet-5.5`. List the IDs your plan exposes with `GET https://apiclaw.biz/v1/models`.
- When the daily request quota is used up the API returns HTTP 429; `*apiclaw.APIError` carries `StatusCode` and `RetryAfter`.
- Already using [openai-go](https://github.com/openai/openai-go) or [go-openai](https://github.com/sashabaranov/go-openai)? Just set the base URL to `https://apiclaw.biz/v1`; this package is only here to show the raw request.
- Docs: [apiclaw.biz/docs](https://apiclaw.biz/docs/) · Support: [t.me/Apiclaw_support](https://t.me/Apiclaw_support)

APIClaw routes requests to models available through publicly available sources and provider APIs. Maintained by the APIClaw team; APIClaw is an independent gateway and is not affiliated with Anthropic, OpenAI or Google. MIT licensed.
