"use client"

import { useState } from "react";
import { useForm, SubmitHandler } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";


import { authClient } from "@/packages/auth/auth-client"
import { handymanSignupRequest, handymanSignupSchema } from "@/lib/schemas/auth.schema";

export function HandymanSignUpForm () {
    
    const [showPassword, setShowPassword] = useState<boolean>(false);
    
    const { register, watch, handleSubmit, formState: { errors, isSubmitting } } = useForm<handymanSignupRequest>({
        resolver: zodResolver(handymanSignupSchema)
    });

    const onSubmit: SubmitHandler<handymanSignupRequest> = async (data) => {
        try{
            const response = await fetch(
                `${process.env.NEXT_PUBLIC_API_URL}/handyman/sign-up`,
                {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    credentials: "include", // important for cookies/session
                    body: JSON.stringify(data),
                }
            );
            
            const result = await response.json();

            if (!response.ok) {
                console.log(result);
                alert(result.error || "Sign Up failed");
                return;
            }

            console.log("Signed in:", result);
            alert("Sign Up Successful, Redirecting");
            
            await authClient.signIn.email({
                email: data.email,
                password: data.password,
            });
            window.location.assign("/") ;
        } catch (error) {
            console.error(error);
            alert("Something went wrong");
        }
    }

    const password = watch("password");

    return(
        <form 
            className="flex flex-col text-black text-xl gap-5"
            onSubmit={ handleSubmit(onSubmit) }
        >
            <div className="flex flex-col gap-1 px-3">
                <label 
                htmlFor="name"
                className="font-semibold"
                >
                    Name
                </label>
                <input 
                    className="border-2 border-black rounded-lg px-2 py-1" 
                    { ...register("name") } type="name"
                    placeholder="Full Name"
                />
                {errors.name && <p>{errors.name.message}</p>}
            </div>
            
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

            <div className="px-3">
                <select 
                    className="border-2 border-black rounded-lg px-15 py-2"
                    { ...register("category")}
                >
                    <option value="">Select Category</option>
                    {
                        ["plumber", "electrician", "carpenter", "mason", "mechanic", "hvac technician", "landscaper", "deep_cleaner"]
                        .map(c=> (<option key={c} value={c}>{c.replace("_", "")}</option>))
                    }
                </select>
                {errors.category && <p>{errors.category.message}</p>}
            </div>
            

            <button
                className="bg-red-500 hover:bg-green-500 transition-colors duration-300 text-white rounded-xl py-2 px-6 font-bold w-fit self-center"
                type="submit"
                disabled={isSubmitting}
            >
                {isSubmitting ? "Signing Up..." : "Sign Up"}
            </button>
        </form>
    )
} 