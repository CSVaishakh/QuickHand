import type { Session } from "better-auth";

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
    category: "plumber" | "electrician" | "carpenter" | "mason" | "mechanic" | "havc_technician" | "landscaper" | "cleaner",
    id: string 
}

export type Variables = {
    session: Session,
    user : user,
    jobs: any[]
}

export type listed_job = {
    id: string,
    name: string,
    customer: string,
    pay_range: string,
    job_category: "plumbing" | "electrical" | "carpentery" | "masonary" | "mechanical" | "havc" | "landscaping" | "cleaning"
}