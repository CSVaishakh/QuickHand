import { eq } from "drizzle-orm";
import { db } from "../db";
import { user } from "../db/schema";
import { type signupRequest, type signupResponse } from "../lib/types";
import { Hono } from "hono";

const customerRouter = new Hono();

customerRouter.post('/sign-up', async (c) => {
    const body: signupRequest = await c.req.json();
    const { email, password, name} = body;
    
    const origin = new URL(c.req.url).origin;
    const response  = await fetch( `${origin}/auth/sign-up/email`, {
        method: "POST",
        headers: {
            'Content-Type' : "application/json",
        },
        body: JSON.stringify({
            email,
            password,
            name
        })
    });

    const data = await response.json() as {token: string, user: signupResponse};
    const id  = data.user.id;
    const token = data.token;

    if(data) {
        const [updatedUser] = await db
            .update(user)
            .set({ role: 'customer'})
            .where(eq(user.id, id))
            .returning()

        return c.json({
            token: data.token,
            user : updatedUser
        })
    }

    return c.json({ });
})

export default customerRouter;