import { relations } from "drizzle-orm";
import { pgTable, text, integer, timestamp, index } from "drizzle-orm/pg-core";
import { user } from "./auth-schema";

export const jobs = pgTable("jobs", {
    id: text("id").primaryKey(),
    name: text("name").notNull(),
    customer: text("customer")
      .notNull()
      .references(() => user.id, { onDelete: "cascade" }),
    handyman: text("handyman")
      .notNull()
      .references(() => user.id, { onDelete: "cascade" }),
    hired_at: timestamp("hired_at").defaultNow().notNull(),
    cost: integer("cost").notNull(),
    job_status: text("job_status").notNull().$type<"NotCompleted" | "Completed">()
}, (table) => [
  index("jobs_customer_idx").on(table.customer),
  index("jobs_handyman_idx").on(table.handyman),
  index("jobs_status_idx").on(table.job_status)
]);

export const jobRelations = relations(jobs, ({ one }) => ({
  customerUser: one(user, {
    fields: [jobs.customer],
    references: [user.id],
    relationName: "customerJobs"
  }),
  handymanUser: one(user, {
    fields: [jobs.handyman],
    references: [user.id],
    relationName: "handymanJobs"
  })
}));