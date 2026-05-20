    import { betterAuth } from "better-auth";
    import { drizzleAdapter } from "better-auth/adapters/drizzle";
    import { db } from "@repo/db";

    export const  auth = betterAuth({
        secret: process.env.BETTER_AUTH_SECRET,

        database: drizzleAdapter(db,{
            provider: "pg",
        }),

        baseURL: process.env.BETTER_AUTH_URL,
        
        emailAndPassword: {
            enabled: true,
            autoSignIn: true,  
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

