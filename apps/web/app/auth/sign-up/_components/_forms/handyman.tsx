"use client"

import { useForm, SubmitHandler } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { handymanSignupRequest, handymanSignupSchema } from "@/lib/schemas/auth.schema";
import { useRouter } from "next/navigation";

export function HandymanSignUpForm () {
    const { register, handleSubmit, formState: { errors, isSubmitting } } = useForm<handymanSignupRequest>({
        resolver: zodResolver(handymanSignupSchema)
    });

    const router = useRouter();

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
            router.push("/")

        } catch (error) {
            console.error(error);
            alert("Something went wrong");
        }

    }

    return(
        <form 
            className="flex flex-col text-black gap-5 font-semibold"
            onSubmit={ handleSubmit(onSubmit) }
        >
            <div className="px-3">
                <h1>Name</h1>
                <input 
                    className="border-2 border-black rounded-lg px-2" 
                    { ...register("name") } type="name"
                    placeholder="Full Name"
                />
                {errors.name && <p>{errors.name.message}</p>}
            </div>
            
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
                    placeholder="example@123"
                />
                {errors.password && <p>{errors.password.message}</p>}
            </div>

            <div className="px-3">
                <select 
                    className="border-2 border-black rounded-lg px-9.5 py-1"
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