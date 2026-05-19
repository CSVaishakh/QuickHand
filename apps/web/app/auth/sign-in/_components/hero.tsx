"use client"

import { SignInFrom } from "./_forms/form";

export default function AuthHero () {
    return(
        <section className="flex flex-col items-center w-full pt-6">
            <SignInFrom/>
        </section>
    )
}