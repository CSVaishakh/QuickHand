import { betterAuth, string } from "better-auth";
import { drizzleAdapter } from "better-auth/adapters/drizzle";
import { db } from "../../db";

export const  auth = betterAuth({
    database: drizzleAdapter(db,{
        provider: "pg",
    }),
    emailAndPassword: {
        enabled: true,
    },
    user: {
        additionalFields: {
            role: {
                type: "string",
                required: true,
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
});

