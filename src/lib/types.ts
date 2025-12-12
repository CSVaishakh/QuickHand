export type signupRequest = {
    name: string,
    email: string, 
    password: string
    image: string,
}

export type signupResponse = {
    name: string,
    email: string
    emailVerified: boolean ,
    image: string | null,
    createdAt: Date,
    updatedAt: Date,
    role: string | null,
    category: string | null,
    id: string 
}
