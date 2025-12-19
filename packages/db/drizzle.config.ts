import { defineConfig } from "drizzle-kit";

export default defineConfig({
    schema: '/home/vaishakh/Projects/QuickHand/packages/db/schema.ts',
    out: '../../drizzle',
    dialect: 'postgresql',
    dbCredentials: {
        url: process.env.DATABASE_URL!,
    },
});