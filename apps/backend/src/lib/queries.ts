import { db, jobs, listed_jobs, user, eq } from "@repo/db";
import type { accepct_job, job, listed_job } from "./types/types";

export const findRecordsInJobs = async (userType: "customer" | "handyman", userId: string) => {
    return await db
            .select()
            .from(jobs)
            .where(eq(jobs[userType], userId))
}

export const findHandyman = async (userId: string) => {
    return await db
            .select()
            .from(user)
            .where(eq(user.id, userId))
}

export const listJob = async (job: listed_job) => {
    return await db
            .insert(listed_jobs)
            .values(job)
            .returning()
}

export const accepctJob = async (job: accepct_job) => {
    return await db
            .insert(jobs)
            .values(job)
            .returning()
}

export const fetchJobs = async  (job_category: "plumbing" | "electrical" | "carpentery" | "masonary" | "mechanical" | "havc" | "landscaping" | "deep-cleaning" ) => {
    return await db
        .select()
        .from(listed_jobs)
        .where(eq(listed_jobs.job_category, job_category))
}

export const fetchHandymen = async () => {
    return await db
            .select()
            .from(user)
            .where(eq(user.role, "handyman"))
}