# curl + APIClaw

```bash
curl https://apiclaw.biz/v1/chat/completions \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"model":"YOUR_MODEL_ID","messages":[{"role":"user","content":"Hello"}]}'
```

List the models your key can use: `curl https://apiclaw.biz/v1/models -H "Authorization: Bearer YOUR_API_KEY"`
