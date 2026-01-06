import type { Session } from "@repo/auth";

export type signupRequest = {
    name: string,
    email: string, 
    password: string
    image: string,
}

export type user = {
    name: string,
    email: string
    emailVerified: boolean ,
    image: string | null,
    createdAt: Date,
    updatedAt: Date,
    role: "customer" | "handyman",
    category: "plumber" | "electrician" | "carpenter" | "mason" | "mechanic" | "havc_technician" | "landscaper" | "deep_cleaner",
    id: string 
}

export type listed_job = {
    id: string,
    name: string,
    customer: string,
    pay_range: string,
    job_category: "plumbing" | "electrical" | "carpentery" | "masonary" | "mechanical" | "havc" | "landscaping" | "deep_cleaning"
}

export type job = {
    id: string,
    name: string,
    customer: string,
    handyman: string,
    hired_at: Date,
    cost: number,
    job_status: "NotCompleted" | "Completed",
    list_id: string,
    job_category: "plumbing" | "electrical" | "carpentery" | "masonary" | "mechanical" | "havc" | "landscaping" | "deep_cleaning"
}

export type accepct_job = Omit<job, "hired_at">;

export type Variables = {
    session: Session,
    user : user,
    jobs: job[]
}