import { serve } from "bun";
import { Hono } from "hono";
import { cors } from "hono/cors";
import { auth } from "@repo/auth";
import customerRouter from "./src/routes/customer.router";
import handymanRouter from "./src/routes/handyman.router";

const app = new Hono();

app.use(
    "*",
    cors({
        origin: ["http://localhost:3000", "http://localhost:3001"],
        allowMethods: ["GET", "POST", "OPTIONS"],
        allowHeaders: ["Content-Type", "Authorization"],
    })
);

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