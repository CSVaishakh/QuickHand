import { betterAuth } from "better-auth";
import { drizzleAdapter } from "better-auth/adapters/drizzle";
import { db } from "../../db";

export const  auth = betterAuth({
    database: drizzleAdapter(db,{
        provider: "pg",
    }),
    emailAndPassword: {
        enabled: true,
        autoSignIn: false,  
    },
    user: {
        additionalFields: {
            role: {
                type: "string",
                required: false,
                input: false,
            },
            category: {
                type: "string",
                required: false,
                input: false,
                defaultValue: null,
            },
        },
    },
    basePath: "/auth",
    advanced: {
        disableOriginCheck: true
    },
});

