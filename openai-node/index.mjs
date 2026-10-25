// npm install openai
// APICLAW_API_KEY=... APICLAW_MODEL=... node index.mjs
import OpenAI from "openai";

const client = new OpenAI({
  baseURL: "https://apiclaw.biz/v1",
  apiKey: process.env.APICLAW_API_KEY,
});

const resp = await client.chat.completions.create({
  model: process.env.APICLAW_MODEL ?? "YOUR_MODEL_ID",
  messages: [{ role: "user", content: "Say hello in one sentence." }],
});
console.log(resp.choices[0].message.content);
