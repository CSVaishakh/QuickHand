import { auth } from "./auth"
import { createMiddleware } from "hono/factory";

export const requireAuth = createMiddleware(async (c, next) => {
    const sessionData = await auth.api.getSession({ headers: c.req.raw.headers });
    
    if (!sessionData) {
         return c.json({error: "Unauthorized"}, 401)
    }
    
    c.set("session", sessionData);
    c.set("user", sessionData.user);

    await next();
})