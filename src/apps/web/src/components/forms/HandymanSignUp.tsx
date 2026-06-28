import { useForm } from "react-hook-form"
import type { SubmitHandler } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { useNavigate } from "@tanstack/react-router"

import { HandymanType } from "#/types/types"
import type { HandymanSignUpReq } from "#/types/types"
import {
   HandymanSignUpReqSchema,
   HandymanSignUpResSchema,
} from "#/schemas/schemas"

import { Button } from "@components/ui/button"
import { authClient } from "#/auth-client"

import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"

function HandymanSignUpForm() {
   const navigate = useNavigate()

   const {
      register,
      handleSubmit,
      watch,
      setValue,
      formState: { errors },
   } = useForm<HandymanSignUpReq>({
      resolver: zodResolver(HandymanSignUpReqSchema),
   })

   const onSubmit: SubmitHandler<HandymanSignUpReq> = async (data) => {
      try {
         const res = await authClient.HandymanSignUp(data)
         const json = await res.json()

         const parsed = HandymanSignUpResSchema.safeParse(json)

         if (!parsed.success) {
            console.error(parsed.error.issues)
            throw new Error("Invalid API response")
         }

         const res_data = parsed.data

         localStorage.setItem("token", res_data.token)
         localStorage.setItem("user", JSON.stringify(res_data))

         navigate({ to: "/" })
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
                  className="border-2 border-black text-black bg-white"
               />
               <div>
                  {errors.firstName && (
                     <p className="boder-2 border-black bg-red-500 text-white">
                        {errors.firstName.message}
                     </p>
                  )}
               </div>
            </div>

            <div>
               <input
                  {...register("lastName")}
                  placeholder="Last Name"
                  className="border-2 border-black text-black bg-white"
               />
               <div>
                  {errors.lastName && (
                     <p className="boder-2 border-black bg-red-500 text-white">
                        {errors.lastName.message}
                     </p>
                  )}
               </div>
            </div>

            <div>
               <input
                  {...register("email")}
                  type="email"
                  placeholder="Email"
                  className="border-2 border-black text-black bg-white"
               />
               <div>
                  {errors.email && (
                     <p className="boder-2 border-black bg-red-500 text-white">
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
                  className="border-2 border-black text-black bg-white"
               />
               <div>
                  {errors.password && (
                     <p className="boder-2 border-black bg-red-500 text-white">
                        {errors.password.message}
                     </p>
                  )}
               </div> 
            </div>

            <div>
               <input type="hidden" {...register("type")} />
               <Select
                 value={watch("type")}
                 onValueChange={(value) =>
                   setValue("type", value as HandymanType, {
                     shouldValidate: true,
                     shouldDirty: true,
                   })
                 }
               >
                 <SelectTrigger className="w-53 border-2 border-black">
                   <SelectValue placeholder="Select Handyman Type" />
                 </SelectTrigger>
               
                 <SelectContent className="border-2 border-black">
                   <SelectGroup>
                     {Object.values(HandymanType).map((type) => (
                       <SelectItem key={type} value={type}>
                         {type}
                       </SelectItem>
                     ))}
                   </SelectGroup>
                 </SelectContent>
               </Select>

               <div>
                  {errors.type && (
                     <p className="boder-2 border-black bg-red-500 text-white">
                        {errors.type.message}
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

export default HandymanSignUpForm