"use client"

import { useState } from "react";
import { HandymanSignUpForm } from "./_forms/handyman";
import { CustomerSignUpForm } from "./_forms/customer";

export default function SignUpHero () {
    const [choice, setChoice] = useState<"customer" | "handyman">("handyman")
    return(
        <section className="flex flex-col items-center justify-center">
            <div className="flex text-black font-semibold p-4 gap-3">
                <button 
                    className="border-2 border-balck rounded-xl px-5 py-2 hover:text-blue-400"
                    onClick={ (() => setChoice("customer")) }
                >
                    Customer
                </button>
                <button
                    className="border-2 border-balck rounded-xl px-5 py-2 hover:text-blue-400"
                    onClick={ (()=>setChoice("handyman")) }
                >
                    Handyman
                </button>
            </div>

            {choice === "customer" && <CustomerSignUpForm/>}
            {choice === "handyman" && <HandymanSignUpForm/>}
        </section>
    )
}