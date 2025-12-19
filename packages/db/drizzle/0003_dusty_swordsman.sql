CREATE TABLE "listed-jobs" (
	"id" text PRIMARY KEY NOT NULL,
	"name" text NOT NULL,
	"customer" text NOT NULL,
	"pay-range" text NOT NULL
);
--> statement-breakpoint
ALTER TABLE "listed-jobs" ADD CONSTRAINT "listed-jobs_customer_user_id_fk" FOREIGN KEY ("customer") REFERENCES "public"."user"("id") ON DELETE cascade ON UPDATE no action;