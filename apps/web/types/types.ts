export type handyman = {
    name: string,
    email: string,
    password: string,
    category: "plumber" | "electrician" | "carpenter" | "mason" | "mechanic" | "havc_technician" | "landscaper" | "deep_cleaner"
}

export type customer = {
    name: string,
    email: string,
    password: string
}