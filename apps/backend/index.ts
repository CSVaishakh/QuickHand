import { serve } from "bun";
import { Hono } from "hono";
import { auth } from "@repo/auth";
import customerRouter from "./src/routes/customer.router";
import handymanRouter from "./src/routes/handyman.router";

const app = new Hono();

app.on(["POST", "GET"], "/auth/*", (c) => auth.handler(c.req.raw));

app.get('/test', (c) => c.text('Hello World!Server is working'))

app.route('/customer', customerRouter)
app.route('/handyman', handymanRouter)

const port = process.env.PORT

serve({
    fetch: app.fetch,
    port,
});

console.log(`App running on port ${port}`);