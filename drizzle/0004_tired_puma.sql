CREATE TABLE "listed_jobs" (
	"id" text PRIMARY KEY NOT NULL,
	"name" text NOT NULL,
	"customer" text NOT NULL,
	"pay_range" text NOT NULL,
	"created_at" timestamp DEFAULT now() NOT NULL
);
--> statement-breakpoint
ALTER TABLE "listed-jobs" DISABLE ROW LEVEL SECURITY;--> statement-breakpoint
DROP TABLE "listed-jobs" CASCADE;--> statement-breakpoint
ALTER TABLE "jobs" ADD COLUMN "list_id" text NOT NULL;--> statement-breakpoint
ALTER TABLE "listed_jobs" ADD CONSTRAINT "listed_jobs_customer_user_id_fk" FOREIGN KEY ("customer") REFERENCES "public"."user"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
CREATE INDEX "list_jobs_customer_idx" ON "listed_jobs" USING btree ("customer");--> statement-breakpoint
ALTER TABLE "jobs" ADD CONSTRAINT "jobs_list_id_listed_jobs_id_fk" FOREIGN KEY ("list_id") REFERENCES "public"."listed_jobs"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
CREATE INDEX "jobs_list_id_idx" ON "jobs" USING btree ("list_id");