export enum UserRole {
    HANDYMAN = "handyman",
    CUSTOMER = "client",
}

export enum HandymanType {
   PLUMBER           = "plumber",
   ELECTRICIAN       = "electrician",
   CARPENTER         = "carpenter",
   MASON             = "mason",
   MECHANIC          = "mechanic",
   HVAC_TECHNICIAN   = "hvac_technician",
   LANDSCAPER        = "landscaper",
   DEEP_CLEANER      = "deep_cleaner",
}

export interface SignInReq {
    email:     string
    password:  string
    role:      UserRole
}

export interface HandymanSignUpReq {
    firstName:    string
    lastName:     string
    phone_number: string
    email:        string
    password:     string
    type:         HandymanType
}

export interface ClientSignUpReq {
    firstName:    string
    lastName:     string
    phone_number: string
    email:        string
    password:     string
}

export interface GetSessionReq {
   token: string
}

export interface HandymanSignInRes {
   userId:     string
   firstName:  string
   token: 		string
	role:			UserRole
	type:			HandymanType
}

export interface ClientSignInRes {
   userId:     string
   firstName:  string
   token: 		string
	role:			UserRole
}

export interface HandymanSignUpRes {
   userId:     string
   firstName:  string
   token: 		string
   role:       UserRole
	type:       HandymanType
}

export interface ClientSignUpRes {
   userId:     string
   firstName:  string
   token: 		string
   role:       UserRole
}

export interface GetSessionRes {
   sessionId:     string
   revoked:       boolean
   createdAt:     string
   userId:        string
   firstName:     string
   email:         string
   role:          UserRole
   type:          HandymanType | null
}