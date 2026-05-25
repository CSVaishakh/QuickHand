"use client"

import { useForm, SubmitHandler } from "react-hook-form";
import{ zodResolver } from "@hookform/resolvers/zod";
import { authClient } from "@/packages/auth/auth-client";

import { SigninRequest, SigninSchema } from "@/lib/schemas/auth.schema";

export function SignInFrom () {
    const { register, handleSubmit, formState: { errors, isSubmitting }} = useForm<SigninRequest>({
        resolver: zodResolver(SigninSchema)
    });

    const onSubmit: SubmitHandler<SigninRequest> = async (data) => {
        try {

        await authClient.signIn.email({
            email: data.email,
            password: data.password,
        });

        alert("Sign in successful");

        window.location.assign("/");

        } catch (error) {
            console.error(error);
            alert("Sign in failed");
        }
    };
        
    return(
        <form
            className="flex flex-col text-black gap-5" 
            onSubmit={ handleSubmit(onSubmit) }
        >
            <div className="px-3">
                <h1>Email</h1>
                <input 
                    className="border-2 border-black rounded-lg px-2" 
                    {...register("email")} 
                    type="email" 
                    placeholder="example@gmail.com"
                />
                {errors.email && <p>{errors.email.message}</p>}
            </div>

            <div className="px-3">
                <h1>Password</h1>
                <input 
                    className="border-2 border-black rounded-lg px-2" 
                    {...register("password")} 
                    type="password" 
                    placeholder="exam
                    ple@123"
                />
                {errors.password && <p>{errors.password.message}</p>}
            </div>

            <button
                className="bg-red-500 hover:bg-green-500 transition-colors duration-300 text-white rounded-xl py-2 px-6 font-semibold w-fit self-center"
                type="submit"
                disabled={isSubmitting}
                >
                    {isSubmitting ? "Signing In..." : "Sign In"}
            </button>
        </form>
    )
}