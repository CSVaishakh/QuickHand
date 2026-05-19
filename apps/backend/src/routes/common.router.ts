import { Hono } from "hono";
import type { Variables } from "../lib/types/types";
import { SigninSchema } from "../lib/schemas/common.schema";
import { z } from "zod";

const commonRouter = new Hono<{Variables: Variables}>();

commonRouter.post('/sign-in', async (c) => {
    const raw = await c.req.json();

    const result = SigninSchema.safeParse(raw)
    if(!result.success){
        return c.json({ error: z.treeifyError(result.error) }, 400)
    }
    const { email, password } = result.data;

    const origin = new URL(c.req.url).origin;
    const response = await fetch(`${origin}/auth/sign-in/email`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
    });

    if (!response.ok) {
        const error = await response.json();
        return c.json({ error: 'Sign in failed', details: error }, response.status as 400 | 401 | 500);
    }

    const setCookie = response.headers.get('set-cookie');
    if (setCookie) {
        c.header('Set-Cookie', setCookie);
    }

    const data = await response.json();
    return c.json(data);
});
export default commonRouter;