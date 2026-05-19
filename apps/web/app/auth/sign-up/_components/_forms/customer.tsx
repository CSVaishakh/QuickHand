import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { customerSignupRequest, customerSignupSchema } from "@/lib/schemas/auth.schema";

export function CustomerSignUpForm () {
    const { register, handleSubmit, formState: { errors } } = useForm<customerSignupRequest>({
        resolver: zodResolver(customerSignupSchema)
    });

    return(
        <form onSubmit={ handleSubmit(data => console.log(data)) }>
            <input { ...register("name") } type="name" />
            {errors.name && <p>{errors.name.message}</p>}

            <input {...register("email") } type="email"/>
            {errors.email && <p>{errors.email.message}</p>}

            <input {...register("password")} type="password"/>
            {errors.password && <p>{errors.password.message}</p>}
        </form>
    )
}