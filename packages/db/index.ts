import { drizzle } from "drizzle-orm/postgres-js";
import postgres from "postgres";
import * as schema from "./schema";

export * from "./schema";
export * from "drizzle-orm";

const db_client = postgres(process.env.DATABASE_URL!);
export const db = drizzle(db_client, { schema });