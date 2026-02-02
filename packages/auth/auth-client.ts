import { createAuthClient } from "better-auth/react";

export const authClient = createAuthClient({
    baseURL: process.env.NEXT_PUBLIC_WEB_APP_URL
});

export const { getSession, useSession } = authClient;
