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
   email:      string
   passowrd:   string
   role:       UserRole
}

export interface HandymanSignUpReq{
   firstName:    string
   lastName:     string
   email:        string
   password:     string
   type:         HandymanType
}

export interface ClientSignUpReq{
   firstName:    string
   lastName:     string
   email:        string
   password:     string
}