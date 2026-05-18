import { z } from "zod";

export const customerSignupSchema = z.object({
    name: z.string().min(1),
    email: z.string().min(1),
    password: z.string().min(8),
    image: z.string().pipe(z.url()).optional()
})

export const customerSigninSchema = z.object({
    email: z.string().min(1),
    password: z.string().min(8)
})

export const customerListedjobSchema = z.object({
    id: z.string(),
    name: z.string().min(1),
    customer: z.string().min(1),
    pay_range: z.string(),
    job_category: z.enum(["plumbing", "electrical", "carpentery", "masonary", "mechanical", "havc", "landscaping", "deep_cleaning"])
})

export type customerSignupRequest = z.infer<typeof customerSignupSchema>;
export type customerSigninRequest = z.infer<typeof customerSigninSchema>;
export type customerListedjobRequest = z.infer<typeof customerListedjobSchema>;