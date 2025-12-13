import { auth } from "./auth"
import { createMiddleware } from "hono/factory";

export const requireAuth = createMiddleware(async (c, next) => {
    const sessionData = await auth.api.getSession({ headers: c.req.raw.headers });
    
    if (sessionData) {
        c.set("Session", sessionData);
        await next();
        return
    }else{
         return c.json({error: "Unauthorized"}, 401)
    }
})