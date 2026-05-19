import { z } from "zod";

export const customerSignupSchema = z.object({
    name: z.string().min(1),
    email: z.email(),
    password: z.string().min(8),
    image: z.url().optional()
})

export const customerListedjobSchema = z.object({
    id: z.string(),
    name: z.string().min(1),
    customer: z.string().min(1),
    pay_range: z.string(),
    job_category: z.enum(["plumbing", "electrical", "carpentery", "masonary", "mechanical", "hvac", "landscaping", "deep_cleaning"])
})

export type customerSignupRequest = z.infer<typeof customerSignupSchema>;
export type customerListedjobRequest = z.infer<typeof customerListedjobSchema>;