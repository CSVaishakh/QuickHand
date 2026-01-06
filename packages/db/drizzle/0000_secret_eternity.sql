CREATE TABLE "account" (
	"id" text PRIMARY KEY NOT NULL,
	"account_id" text NOT NULL,
	"provider_id" text NOT NULL,
	"user_id" text NOT NULL,
	"access_token" text,
	"refresh_token" text,
	"id_token" text,
	"access_token_expires_at" timestamp,
	"refresh_token_expires_at" timestamp,
	"scope" text,
	"password" text,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp NOT NULL
);
--> statement-breakpoint
CREATE TABLE "jobs" (
	"id" text PRIMARY KEY NOT NULL,
	"name" text NOT NULL,
	"customer" text NOT NULL,
	"handyman" text NOT NULL,
	"hired_at" timestamp DEFAULT now() NOT NULL,
	"cost" integer NOT NULL,
	"job_status" text NOT NULL,
	"list_id" text NOT NULL,
	"worker_category" text NOT NULL
);
--> statement-breakpoint
CREATE TABLE "listed_jobs" (
	"id" text PRIMARY KEY NOT NULL,
	"name" text NOT NULL,
	"customer" text NOT NULL,
	"pay_range" text NOT NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"job_category" text NOT NULL
);
--> statement-breakpoint
CREATE TABLE "session" (
	"id" text PRIMARY KEY NOT NULL,
	"expires_at" timestamp NOT NULL,
	"token" text NOT NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp NOT NULL,
	"ip_address" text,
	"user_agent" text,
	"user_id" text NOT NULL,
	CONSTRAINT "session_token_unique" UNIQUE("token")
);
--> statement-breakpoint
CREATE TABLE "user" (
	"id" text PRIMARY KEY NOT NULL,
	"name" text NOT NULL,
	"email" text NOT NULL,
	"email_verified" boolean DEFAULT false NOT NULL,
	"image" text,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL,
	"role" text,
	"category" text,
	CONSTRAINT "user_email_unique" UNIQUE("email")
);
--> statement-breakpoint
CREATE TABLE "verification" (
	"id" text PRIMARY KEY NOT NULL,
	"identifier" text NOT NULL,
	"value" text NOT NULL,
	"expires_at" timestamp NOT NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL
);
--> statement-breakpoint
ALTER TABLE "account" ADD CONSTRAINT "account_user_id_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."user"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "jobs" ADD CONSTRAINT "jobs_customer_user_id_fk" FOREIGN KEY ("customer") REFERENCES "public"."user"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "jobs" ADD CONSTRAINT "jobs_handyman_user_id_fk" FOREIGN KEY ("handyman") REFERENCES "public"."user"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "jobs" ADD CONSTRAINT "jobs_list_id_listed_jobs_id_fk" FOREIGN KEY ("list_id") REFERENCES "public"."listed_jobs"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "listed_jobs" ADD CONSTRAINT "listed_jobs_customer_user_id_fk" FOREIGN KEY ("customer") REFERENCES "public"."user"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "session" ADD CONSTRAINT "session_user_id_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."user"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
CREATE INDEX "account_userId_idx" ON "account" USING btree ("user_id");--> statement-breakpoint
CREATE INDEX "jobs_customer_idx" ON "jobs" USING btree ("customer");--> statement-breakpoint
CREATE INDEX "jobs_handyman_idx" ON "jobs" USING btree ("handyman");--> statement-breakpoint
CREATE INDEX "jobs_status_idx" ON "jobs" USING btree ("job_status");--> statement-breakpoint
CREATE INDEX "jobs_list_id_idx" ON "jobs" USING btree ("list_id");--> statement-breakpoint
CREATE INDEX "listed_jobs_customer_idx" ON "listed_jobs" USING btree ("customer");--> statement-breakpoint
CREATE INDEX "session_userId_idx" ON "session" USING btree ("user_id");--> statement-breakpoint
CREATE INDEX "verification_identifier_idx" ON "verification" USING btree ("identifier");