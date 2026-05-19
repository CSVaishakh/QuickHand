import { z } from "zod";

export const SigninSchema = z.object({
    email: z.email(),
    password: z.string().min(8)
})

export type customerSigninRequest = z.infer<typeof SigninSchema>;