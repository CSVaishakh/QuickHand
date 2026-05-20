import { z } from "zod";

export const SigninSchema = z.object({
    email: z.email(),
    password: z.string().min(8)
})

export const UpdateImageSchema = z.object({
    img: z.url(),
    userId: z.string()
})

export type customerSigninRequest = z.infer<typeof SigninSchema>;
export type updateImageRequest = z.infer<typeof UpdateImageSchema>;