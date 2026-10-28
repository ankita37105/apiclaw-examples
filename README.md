# APIClaw Examples: Flat-Rate OpenAI-Compatible API Configs for Claude Code, Cursor, Cline & More

Copy-paste configs for using **[APIClaw](https://apiclaw.biz)**, a flat-rate, OpenAI-compatible LLM API gateway, with popular coding agents, SDKs and frameworks.

> Maintained by the **APIClaw team**. APIClaw is an independent gateway, not affiliated with Anthropic, OpenAI, Google or Cursor.

## What APIClaw is (30 seconds)

- **One key, many providers:** Claude, GPT, Gemini, Mistral, Llama, Cohere, Grok, DeepSeek, Qwen, Kimi and GLM through one OpenAI-compatible endpoint. List IDs with `GET https://apiclaw.biz/v1/models` or browse the [model hub](https://apiclaw.biz/ui/model_hub/).
- **Flat monthly price, billed by requests, not tokens:** $19/month (500 requests/day) up to $129/month (8,000 requests/day). [Plans](https://apiclaw.biz/).
- **Will it be cheaper for you?** Check with the free [LLM API Cost Calculator for Coding Agents](https://ankita37105.github.io/llm-api-cost-notes/calculator.html). Heavier models count as more than one request, and for light use per-token billing is often cheaper.
- **Free trial:** 50 requests, no card. [Sign up](https://apiclaw.biz/ui/signup/) → [create a key](https://apiclaw.biz/ui/api-keys/).

## Base URLs, the one thing to get right

| Client speaks… | Base URL |
|---|---|
| OpenAI protocol (Cursor, Cline, Continue, Aider, OpenAI SDKs, LangChain, LlamaIndex, Vercel AI SDK, Codex CLI) | `https://apiclaw.biz/v1` |
| Anthropic Messages protocol (Claude Code, Anthropic SDK) | `https://apiclaw.biz` (**no** `/v1`) |
| Tools that want the full endpoint URL (e.g. VS Code / Copilot custom endpoint) | `https://apiclaw.biz/v1/chat/completions` |

The examples default to **`apiclaw/claude-sonnet-5.5`**. Other real model IDs you can drop in:

| Provider | Example IDs |
|---|---|
| Claude | `apiclaw/claude-sonnet-5.5`, `apiclaw/claude-opus-5.5`, `apiclaw/claude-fable-5.1` |
| GPT | `apiclaw/gpt-6-sol` |
| DeepSeek | `apiclaw/deepseek-v4-pro`, `apiclaw/deepseek-v4-flash` |
| Qwen | `apiclaw/qwen3.8-max-0902`, `apiclaw/qwen3.8-flash` |
| Kimi | `kimi-k3` |
| GLM | `apiclaw/glm-5.3`, `apiclaw/glm-5.3-flash` |

Gemini, Mistral, Llama, Cohere and Grok models are also available; run `curl https://apiclaw.biz/v1/models -H "Authorization: Bearer $APICLAW_API_KEY"` to see the exact IDs your plan exposes. Replace `YOUR_API_KEY` / `APICLAW_API_KEY` with your key.

## Examples

| Tool | Folder | Official APIClaw guide |
|---|---|---|
| Claude Code | [claude-code/](claude-code/) | [apiclaw.biz/docs/claude-code/](https://apiclaw.biz/docs/claude-code/) |
| Cursor | [cursor/](cursor/) | [apiclaw.biz/docs/cursor/](https://apiclaw.biz/docs/cursor/) |
| Cline | [cline/](cline/) | [apiclaw.biz/docs/cline/](https://apiclaw.biz/docs/cline/) |
| Continue | [continue/](continue/) | [apiclaw.biz/docs/vscode/](https://apiclaw.biz/docs/vscode/) |
| Codex CLI | [codex-cli/](codex-cli/) | [apiclaw.biz/docs/codex-cli/](https://apiclaw.biz/docs/codex-cli/) |
| Aider | [aider/](aider/) | standard OpenAI-compatible setup |
| OpenAI Python SDK | [openai-python/](openai-python/) | standard OpenAI-compatible setup |
| OpenAI Node SDK | [openai-node/](openai-node/) | standard OpenAI-compatible setup |
| LangChain (Python) | [langchain/](langchain/) | standard OpenAI-compatible setup |
| LlamaIndex | [llamaindex/](llamaindex/) | standard OpenAI-compatible setup |
| Vercel AI SDK | [vercel-ai-sdk/](vercel-ai-sdk/) | standard OpenAI-compatible setup |
| Go (stdlib client + CLI, on [pkg.go.dev](https://pkg.go.dev/github.com/ankita37105/apiclaw-examples/go)) | [go/](go/) | standard OpenAI-compatible setup |
| curl | [curl/](curl/) | [apiclaw.biz/docs/](https://apiclaw.biz/docs/) |

Examples marked "standard OpenAI-compatible setup" use each library's documented custom-base-URL option. They aren't covered by a dedicated APIClaw docs page yet. If something breaks, please [open an issue](../../issues).

## FAQ

**Is there a flat-rate API for Claude Code?** Yes. Set `ANTHROPIC_BASE_URL=https://apiclaw.biz` and your APIClaw key in `~/.claude/settings.json` ([example](claude-code/)). Usage is counted per request against a daily quota that resets at 00:00 UTC.

**Can I use a custom API key in Cursor?** Yes. Use Settings → Models → OpenAI API Key, then *Override OpenAI Base URL* = `https://apiclaw.biz/v1`, and add a custom model ([example](cursor/)).

**How does it compare with OpenRouter, Requesty, Portkey, LiteLLM, Together and Groq?** See the sourced comparison: **[openai-compatible-api-gateways](https://github.com/ankita37105/openai-compatible-api-gateways)**.

**When shouldn't I use it?** When you need vendor BAAs or enterprise SLAs, unlimited request volume, card billing, or only make a few small calls a day (per-token is cheaper then).

## License

MIT. See [LICENSE](LICENSE).
