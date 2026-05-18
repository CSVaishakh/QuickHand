export type handyman = {
    name: string,
    email: string,
    password: string,
    category: "plumber" | "electrician" | "carpenter" | "mason" | "mechanic" | "hvac_technician" | "landscaper" | "deep_cleaner"
}

export type customer = {
    name: string,
    email: string,
    password: string
}