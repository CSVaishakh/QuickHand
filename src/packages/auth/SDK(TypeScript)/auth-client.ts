import { Role, type ClientSignUpReq, type HandymanSignUpReq, type SignInRequest } from "./types";

export interface AuthClientConfig {
    baseURL: string;
    baseRoute?: string;
}

export class AuthClient {
    private readonly config: Required<AuthClientConfig>
    constructor (
        config: AuthClientConfig
    ){
        this.config = {
            baseRoute: "/api/auth",
            ...config
        } as Required<AuthClientConfig>
    }

    SignIn(data: SignInRequest){
        var url = ""
        if (data.role === Role.HANDYMAN){
            url = `${this.config.baseURL}${this.config.baseRoute}/handyman/sign-in`
        }else if (data.role == Role.CUSTOMER){
            url = `${this.config.baseURL}${this.config.baseRoute}/client/sign-in`
        }
        return fetch(
            url,
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data)
            }
        )
    }
    
    HandymanSignUp(data: HandymanSignUpReq){
        return fetch(
            `${this.config.baseURL}${this.config.baseRoute}/handyman/sign-up`,
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data)
            }
        )
    }
    
    CustomerSignUp(data: ClientSignUpReq){
        return fetch(
            `${this.config.baseURL}${this.config.baseRoute}/client/sign-up`,
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data)
            }
        )
    }

    SignOut(token: string) {
        return fetch(
            `${this.config.baseURL}${this.config.baseRoute}/sign-out`,
            {
                method: "POST",
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            }
        );
    }
}