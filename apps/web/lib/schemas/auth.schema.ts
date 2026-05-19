import { z } from "zod";

export const SigninSchema = z.object({
    email: z.email(),
    password: z.string().min(8)
})

export const customerSignupSchema = z.object({
    name: z.string().min(1),
    email: z.email(),
    password: z.string().min(8),
    image: z.url().optional()
})

export const handymanSignupSchema = z.object({
  name: z.string().min(1),
  email: z.email(),
  password: z.string().min(8),
  image: z.url().optional(),
  category: z.enum(["plumber", "electrician", "carpenter", "mason", "mechanic", "hvac_technician", "landscaper", "deep_cleaner"]),
});


export type customerSignupRequest = z.infer<typeof customerSignupSchema>;
export type handymanSignupRequest = z.infer<typeof handymanSignupSchema>;
export type SigninRequest = z.infer<typeof SigninSchema>;