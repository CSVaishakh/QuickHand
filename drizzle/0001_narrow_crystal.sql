ALTER TABLE "jobs" ADD CONSTRAINT "jobs_customer_user_id_fk" FOREIGN KEY ("customer") REFERENCES "public"."user"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "jobs" ADD CONSTRAINT "jobs_handyman_user_id_fk" FOREIGN KEY ("handyman") REFERENCES "public"."user"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
CREATE INDEX "jobs_customer_idx" ON "jobs" USING btree ("customer");--> statement-breakpoint
CREATE INDEX "jobs_handyman_idx" ON "jobs" USING btree ("handyman");--> statement-breakpoint
CREATE INDEX "jobs_status_idx" ON "jobs" USING btree ("job_status");