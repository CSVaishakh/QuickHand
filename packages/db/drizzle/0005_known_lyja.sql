DROP INDEX "list_jobs_customer_idx";--> statement-breakpoint
ALTER TABLE "jobs" ADD COLUMN "worker_category" text;--> statement-breakpoint
ALTER TABLE "listed_jobs" ADD COLUMN "job_category" text;--> statement-breakpoint
CREATE INDEX "listed_jobs_customer_idx" ON "listed_jobs" USING btree ("customer");