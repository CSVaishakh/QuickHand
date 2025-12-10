import { serve } from "bun";
import { Hono } from "hono";
import { auth } from "./src/lib/auth";

const app = new Hono();

app.on(["POST", "GET"], "/auth/*", (c) => auth.handler(c.req.raw));

app.get('/', (c) => c.text('Hello World!'))

const port = process.env.PORT

serve({
    fetch: app.fetch,
    port,
});

console.log(`App running on port ${port}`);