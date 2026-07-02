import { useForm } from "react-hook-form"
import type { SubmitHandler } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { useNavigate } from "@tanstack/react-router"

import type { SignInReq } from "#/types/types"
import {
   SignInReqSchema,
   SignInResSchema,
} from "#/schemas/schemas"
import { Button } from "@components/ui/button"

import { authClient } from "#/auth-client"

function SignInForm() {
   const navigate = useNavigate()
   const {
      register,
      handleSubmit,
      formState: { errors },
   } = useForm<SignInReq>({
      resolver: zodResolver(SignInReqSchema),
   })

   const onSubmit: SubmitHandler<SignInReq> = async (data) => {
      try {
         const res = await authClient.SignIn(data)
         const json = await res.json()
         const parsed = SignInResSchema.safeParse(json)
         
         if (!parsed.success) {
            console.error(parsed.error.issues);
            throw new Error("Invalid API response")
         }
         
         const res_data = parsed.data;
         localStorage.setItem('token', res_data.token)
         localStorage.setItem("user", JSON.stringify(res_data))
         navigate({ to: '/' })
      } catch (err) {
         console.error(err)
      }
   }

   return (
      <form onSubmit={handleSubmit(onSubmit)}>
         <div className="flex flex-col gap-4">
            <div>
               <input
                  {...register("email")}
                  placeholder="Email"
                  type="email"
                  className="border-2 border-black rounded-xl text-black bg-white"
               />
               <div>
                  {errors.email && (
                     <p className="boder-2 border-black rounded-xl bg-red-500 text-white">
                        {errors.email.message}
                     </p>
                  )}
               </div>   
            </div>
            
            <div>
               <input
                  {...register("password")}
                  type="password"
                  placeholder="Password"
                  className="border-2 border-black rounded-xl text-black bg-white"
               />
               <div>
                  {errors.password && (
                     <p className="boder-2 border-black rounded-xl bg-red-500 text-white">
                        {errors.password.message}
                     </p>
                  )}
               </div>   
            </div>

            <div>
               <Button
                  className="w-22 border-2 rounded-xl bg-red-500 hover:bg-green-600 text-base p-4"
                  size="lg"
                  type="submit"
               >
                  Sign In
               </Button>
            </div>
         </div>
      </form>
   )
}

export default SignInForm