'use client'

import HandymenForm from "./handymen"
import CustomerForm from "./customer"
import { useState } from "react"


export default function SignUpForm() {
    const [choice,setChoice] = useState<"Customer" | "Handymen">("Customer")
    return(
        <section>
            <div className="flex flex-wrap gap-2">
                <button
                    className={`rounded-full px-4 py-2 text-sm font-semibold border transition ${choice === "Customer" ? "border-blue-400 bg-blue-400 text-black" : "border-blue-300 bg-white text-black"}`}
                        type="button"
                        onClick={() => {setChoice("Customer")}}
                        aria-pressed={choice === "Customer"}
                >
                    Customer
                </button>

                <button
                    className={`rounded-full px-4 py-2 text-sm font-semibold border transition ${choice === "Handymen" ? "border-blue-400 bg-blue-400 text-black" : "border-blue-300 bg-white text-black"}`}
                        type="button"
                        onClick={() => {setChoice("Handymen")}}
                        aria-pressed = {choice === "Handymen"}
                >
                    Handymen
                </button>
            </div>

            <div className="mt-4">
                {choice === "Customer" ? <CustomerForm/> : <HandymenForm/>}
            </div>
        </section>
    )
}