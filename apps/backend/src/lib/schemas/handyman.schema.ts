import { z } from "zod";

export const handymanSignupSchema = z.object({
  name: z.string().min(1),
  email: z.string().email(),
  password: z.string().min(8),
  image: z.string().optional(),
  category: z.enum(["plumber", "electrician", "carpenter", "mason", "mechanic", "havc_technician", "landscaper", "deep_cleaner"]),
});

export const handymanSigninSchema = z.object({
    email: z.string().min(1),
    password: z.string().min(8)
})

export const handymanAccepctjobSchema = z.object({
    list_id: z.string(),
    job_name: z.string().min(1),
    customer: z.string().min(1),
    pay_range: z.string(),
    job_category: z.enum(["plumbing", "electrical", "carpentery", "masonary", "mechanical", "havc", "landscaping", "deep_cleaning"])
})

export type handymenSignupRequest = z.infer<typeof handymanSignupSchema>;
export type handymanSigninRequest = z.infer<typeof handymanSigninSchema>;
export type handymanAccepctjobRequest = z.infer<typeof handymanAccepctjobSchema>;