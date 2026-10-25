# Claude Code + APIClaw (flat-rate Claude API for Claude Code)

1. Get a key at https://apiclaw.biz/ui/api-keys/ (the example uses `apiclaw/claude-sonnet-5.5`; swap in `apiclaw/claude-opus-5.5` or any ID from https://apiclaw.biz/ui/model_hub/).
2. Merge [`settings.json`](settings.json) into `~/.claude/settings.json` (user-wide) or `.claude/settings.json` (per project).
3. Start `claude` and run `/status`. The base URL should show `https://apiclaw.biz`.

Notes:
- Claude Code speaks the Anthropic Messages API, so the base URL is **`https://apiclaw.biz` with no `/v1`**.
- Requests count against your plan's daily quota (resets 00:00 UTC). Tokens per request are not capped.

Official guide: https://apiclaw.biz/docs/claude-code/ · Setup walkthrough: https://apiclaw.biz/claude-code-api-setup/
