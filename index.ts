import { serve } from "bun";
import { Hono } from "hono";

const app = new Hono();

app.get('/', (c) => c.text('Hello World!'))

const port = process.env.PORT

serve({
    fetch: app.fetch,
    port,
});

console.log(`App running on port ${port}`);