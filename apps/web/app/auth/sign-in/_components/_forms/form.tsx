"use client"

import { useForm, SubmitHandler } from "react-hook-form";
import { useRouter } from "next/navigation";
import{ zodResolver } from "@hookform/resolvers/zod";
import { SigninRequest, SigninSchema } from "@/lib/schemas/auth.schema";

export function SignInFrom () {
    const { register, handleSubmit, formState: { errors, isSubmitting }} = useForm<SigninRequest>({
        resolver: zodResolver(SigninSchema)
    });

    const router = useRouter();

    const onSubmit: SubmitHandler<SigninRequest> = async (data) => {
        try {
            const response = await fetch(
                `${process.env.NEXT_PUBLIC_API_URL}/common/sign-in`,
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
                alert(result.error || "Sign in failed");
                return;
            }

            console.log("Signed in:", result);
            alert("Sign in Successful, Redirecting");
            router.push("/")

        } catch (error) {
            console.error(error);
            alert("Something went wrong");
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