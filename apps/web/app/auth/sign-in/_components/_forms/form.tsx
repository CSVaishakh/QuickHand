"use client"

import { useState } from "react";
import { useForm, SubmitHandler } from "react-hook-form";
import{ zodResolver } from "@hookform/resolvers/zod";

import { authClient } from "@/packages/auth/auth-client";
import { SigninRequest, SigninSchema } from "@/lib/schemas/auth.schema";

export function SignInForm () {

    const [showPassword, setShowPassword] = useState<boolean>(false);

    const { register,  watch, handleSubmit, formState: { errors, isSubmitting }} = useForm<SigninRequest>({
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
        
    const password = watch("password");

    return(
        <form
            className="flex flex-col text-black text-xl gap-5" 
            onSubmit={ handleSubmit(onSubmit) }
        >
            <div className="flex flex-col gap-1 px-3">
                <label 
                    htmlFor="email"
                    className="font-semibold"
                >
                    Email
                </label>
                <input 
                    className="border-2 border-black rounded-lg px-2 py-1" 
                    {...register("email")} 
                    type="email" 
                    placeholder="example@gmail.com"
                />
                {errors.email && <p>{errors.email.message}</p>}
            </div>


            <div className="px-3 py-1 gap-1">
                <div className="flex justify-between">
                    <label 
                        htmlFor="password"
                        className="font-semibold"
                    >
                        Password
                    </label>
                    {password !== "" &&
                        ( showPassword ? 
                            (
                                <button
                                    type="button"
                                    onClick={(() => {setShowPassword(false)})}
                                >
                                    Hide Password
                                </button>
                            ):(
                                <button
                                    type="button"
                                    onClick={(() => {setShowPassword(true)})}
                                >
                                    Show Password
                                </button>
                            )
                        )
                    }
                </div>
                <input 
                    className="border-2 border-black rounded-lg px-2 py-1" 
                    {...register("password")} 
                    type={showPassword ? "text" : "password"} 
                    placeholder="example@123"
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