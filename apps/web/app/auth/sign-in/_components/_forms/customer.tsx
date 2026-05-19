import { useForm } from "react-hook-form";
import{ zodResolver } from "@hookform/resolvers/zod";
import { customerSigninRequest, customerSigninSchema } from "@/lib/schemas/auth.schema";

export function CustomerSignInFrom () {
    const { register, handleSubmit, formState: { errors }} = useForm<customerSigninRequest>({
        resolver: zodResolver(customerSigninSchema)
    });

    return(
        <form onSubmit={ handleSubmit(data => console.log(data)) }>
            <input {...register("email")} type="email"/>
            {errors.email && <p>{errors.email.message}</p>}

            <input {...register("password")} type="password"/>
            {errors.password && <p>{errors.password.message}</p>}

            <button type="submit">SignIn</button>
        </form>
    )
}