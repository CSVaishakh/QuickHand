import { relations } from "drizzle-orm";
import { pgTable, text, integer, timestamp, index } from "drizzle-orm/pg-core";
import { user } from "./auth-schema";

export const listed_jobs = pgTable("listed_jobs", {
  id: text('id').primaryKey(),
  name: text("name").notNull(),
  customer: text("customer")
    .notNull()
    .references(() => user.id, { onDelete: "cascade" }),
  pay_range: text("pay_range").notNull(),
  created_at: timestamp("created_at").defaultNow().notNull(),
  job_category: text("job_category").$type<"plumbing" | "electrical" | "carpentery" | "masonary" | "mechanical" | "havc" | "landscaping" | "deep_cleaning">().notNull()
}, (table) => [
  index("listed_jobs_customer_idx").on(table.customer),
]);

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
    job_status: text("job_status").notNull().$type<"NotCompleted" | "Completed">(),
    list_id: text("list_id")
      .notNull()
      .references(() => listed_jobs.id, { onDelete: "cascade" }),
    job_category: text("worker_category").$type<"plumbing" | "electrical" | "carpentery" | "masonary" | "mechanical" | "havc" | "landscaping" | "deep_cleaning">().notNull()
}, (table) => [
  index("jobs_customer_idx").on(table.customer),
  index("jobs_handyman_idx").on(table.handyman),
  index("jobs_status_idx").on(table.job_status),
  index("jobs_list_id_idx").on(table.list_id),
]);

export const listJobsRelations = relations(listed_jobs, ({ one, many }) => ({
  customer: one(user, {
    fields: [listed_jobs.customer],
    references: [user.id],
    relationName: "customerListedJobs"
  }),
  hiredJobs: many(jobs, {
    relationName: "listJobToJobs"
  })
}));

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
  }),
  listedJob: one(listed_jobs, {
    fields: [jobs.list_id],
    references: [listed_jobs.id],
    relationName: "listJobToJobs"
  })
}));