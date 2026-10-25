// npm install ai @ai-sdk/openai-compatible
import { createOpenAICompatible } from "@ai-sdk/openai-compatible";
import { generateText } from "ai";

const apiclaw = createOpenAICompatible({
  name: "apiclaw",
  baseURL: "https://apiclaw.biz/v1",
  apiKey: process.env.APICLAW_API_KEY,
});

const { text } = await generateText({
  model: apiclaw(process.env.APICLAW_MODEL ?? "YOUR_MODEL_ID"),
  prompt: "Write a haiku about TypeScript.",
});
console.log(text);
