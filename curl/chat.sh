#!/usr/bin/env bash
# Usage: APICLAW_API_KEY=... APICLAW_MODEL=... ./chat.sh
curl https://apiclaw.biz/v1/chat/completions \
  -H "Authorization: Bearer ${APICLAW_API_KEY}" \
  -H "Content-Type: application/json" \
  -d "{\"model\":\"${APICLAW_MODEL:-apiclaw/claude-sonnet-5.5}\",\"messages\":[{\"role\":\"user\",\"content\":\"Hello\"}]}"
