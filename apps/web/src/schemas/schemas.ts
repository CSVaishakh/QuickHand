import { z } from "zod";
import { UserRole, HandymanType } from "#/types/types";

const UserRoleSchema = z.enum([
   UserRole.HANDYMAN,
   UserRole.CUSTOMER
])

export const HandymanTypeSchema = z.enum([
  HandymanType.PLUMBER,
  HandymanType.ELECTRICIAN,
  HandymanType.CARPENTER,
  HandymanType.MASON,
  HandymanType.MECHANIC,
  HandymanType.HVAC_TECHNICIAN,
  HandymanType.LANDSCAPER,
  HandymanType.DEEP_CLEANER,
])

export const SignInReqSchema = z.object({
   email:      z.email(),
   password:   z.string().min(8),
   role:       UserRoleSchema
})

export const HandymanSignUpReqSchema = z.object({
   firstName:     z.string().min(1).max(25),
   lastName:      z.string().min(1).max(25),
   phone_number:  z.string().min(1).max(10),
   email:         z.email(),
   password:      z.string().min(8),
   type:         HandymanTypeSchema,
})

export const ClientSignUpReqSchema = z.object({
   firstName:     z.string().min(1).max(25),
   lastName: z.string().min(1).max(25),
   phone_number:  z.string().min(1).max(10),
   email:         z.email(),
   password:      z.string().min(8),
})

export const  GetSessionReq = z.object({
   token: z.string()
})

export const HandymanSignInResSchema = z.object({
   userId:        z.string().min(1),
   firstName:     z.string().min(1).max(25),
   token: 		   z.string().min(1),
   role:			   UserRoleSchema,
   type:          HandymanTypeSchema,
})

export const ClientSignInResSchema = z.object({
   userId:        z.string().min(1),
   firstName:     z.string().min(1).max(25),
   token: 		   z.string().min(1),
	role:			   UserRoleSchema,
})

export const HandymanSignUpResSchema = z.object({
   userId:        z.string().min(1),
   firstName:     z.string().min(1).max(25),
   token: 		   z.string().min(1),
	role:			   UserRoleSchema,
   type:          HandymanTypeSchema,
})

export const ClientSignUpResSchema = z.object({
   userId:        z.string().min(1),
   firstName:     z.string().min(1).max(25),
   token: 		   z.string().min(1),
   role:			   UserRoleSchema,
})

export const GetSessionResSchema = z.object({
   sessionId:    z.uuid(),
   revoked:      z.boolean(),
   createdAt:    z.iso.datetime(),
   userId:       z.uuid(),
   firstName:    z.string(),
   email:        z.email(),
   role:         UserRoleSchema,
   type:         HandymanTypeSchema.nullable(),
})

export const SignInResSchema = z.union([
    ClientSignInResSchema,
    HandymanSignInResSchema,
]);