import { useForm } from "react-hook-form"
import type { SubmitHandler } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { useNavigate } from "@tanstack/react-router"

import type { ClientSignUpReq } from "#/types/types"
import {
   ClientSignUpReqSchema,
   ClientSignUpResSchema,
} from "#/schemas/schemas"
import { Button } from "@components/ui/button"

import { authClient } from "#/auth-client"

function ClientSignUpForm () {
   const navigate = useNavigate()
   const {
      register,
      handleSubmit,
      formState: { errors },
   } = useForm<ClientSignUpReq>({
      resolver: zodResolver(ClientSignUpReqSchema),
   })

   const onSubmit: SubmitHandler<ClientSignUpReq> = async (data) => {
      try {
         const res = await authClient.ClientSignUp(data)
         const json = await res.json()
         const parsed = ClientSignUpResSchema.safeParse(json)

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
                  {...register("firstName")}
                  placeholder="First Name"
                  type="first_name"
                  className="border-2 border-black rounded-xl text-black bg-white"
               />
               <div>
                  {errors.firstName && (
                     <p className="boder-2 border-black rounded-xl bg-red-500 text-white">
                        {errors.firstName.message}
                     </p>
                  )}
               </div>
            </div>
   
            <div>
               <input
                  {...register("lastName")}
                  placeholder="Last Name"
                  type="last_name"
                  className="border-2 border-black rounded-xl text-black bg-white"
               />
               <div>
                  {errors.lastName && (
                     <p className="boder-2 border-black rounded-xl bg-red-500 text-white">
                        {errors.lastName.message}
                     </p>
                  )}
               </div>
            </div>

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
                 Sign Up
               </Button>
            </div>
         </div>
      </form>
   )
}

export default ClientSignUpForm