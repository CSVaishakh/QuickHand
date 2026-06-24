export enum Role {
    HANDYMAN = "handyman",
    CUSTOMER = "client",
}

export enum HandymanType {
   PLUMBER = "plumber",
   ELECTRICIAN = "electrician",
   CARPENTER = "carpenter",
   MASON = "mason",
   MECHANIC = "mechanic",
   HVAC_TECHNICIAN = "hvac_technician",
   LANDSCAPER = "landscaper",
   DEEP_CLEANER = "deep_cleaner",
}

export interface SignInRequest {
    email: string
    passowrd: string
    role: Role
}

export interface HandymanSignUpReq{
    firstName: string
    lastName: string
    email: string
    password: string
    type: HandymanType
}

export interface ClientSignUpReq{
    firstName: string
    lastName: string
    email: string
    password: string
}