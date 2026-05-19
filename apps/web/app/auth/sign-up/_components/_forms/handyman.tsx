import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { handymanSignupRequest, handymanSignupSchema } from "@/lib/schemas/auth.schema";

export function HandymanSignUpForm () {
    const { register, handleSubmit, formState: { errors } } = useForm<handymanSignupRequest>({
        resolver: zodResolver(handymanSignupSchema)
    });

    return(
        <form onSubmit={ handleSubmit(data => console.log(data)) }>
            <input { ...register("name") } type="name" />
            {errors.name && <p>{errors.name.message}</p>}

            <input {...register("email") } type="email"/>
            {errors.email && <p>{errors.email.message}</p>}

            <input {...register("password")} type="password"/>
            {errors.password && <p>{errors.password.message}</p>}

            <select{ ...register("category") }>
                <option value="">Select Category</option>
                {["plumber", "electrician", "carpenter", "mason", "mechanic", "hvac_technician", "landscaper", "deep_cleaner"].map( c=> (
                    <option key={c} value={c}>{c.replace("_", "")}</option>
                ))}
            </select>
            {errors.category && <p>{errors.category.message}</p>}
        </form>
    )
} 