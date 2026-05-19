import { useForm } from "react-hook-form";
import{ zodResolver } from "@hookform/resolvers/zod";
import { SigninRequest, SigninSchema } from "@/lib/schemas/auth.schema";

export function SignInFrom () {
    const { register, handleSubmit, formState: { errors }} = useForm<SigninRequest>({
        resolver: zodResolver(SigninSchema)
    });

    return(
        <form
            className="flex flex-col text-black gap-5" 
            onSubmit={ handleSubmit(data => console.log(data)) }
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
                        placeholder="example@123"
                    />
                    {errors.password && <p>{errors.password.message}</p>}
                </div>

                <button
                    className="bg-red-500 hover:bg-green-500 transition-colors duration-300 text-white rounded-xl py-2 px-6 font-semibold w-fit self-center"
                    type="submit"
                    >
                    Sign In
                </button>
        </form>
    )
}