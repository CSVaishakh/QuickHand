"use client"

import { useEffect, useState } from "react";
import SignUpForm from "./_components/signup/signupform";
import SigninForm from "./_components/signin/signinForm";
import { useRouter, useSearchParams } from "next/navigation";

const getChoiceFromMode = (mode: string | null): "Signin" | "Signup" => {
    return mode === "signup" ? "Signup" : "Signin";
};

export default function Auth() {
    const router = useRouter();
    const searchParams = useSearchParams();
    const [choice, setChoice] = useState<"Signin" | "Signup">(
        getChoiceFromMode(searchParams.get("mode"))
    );

    useEffect(() => {
        setChoice(getChoiceFromMode(searchParams.get("mode")));
    }, [searchParams]);

    const handleChoiceChange = (nextChoice: "Signin" | "Signup") => {
        setChoice(nextChoice);
        const params = new URLSearchParams(searchParams.toString());
        params.set("mode", nextChoice === "Signup" ? "signup" : "signin");
        router.replace(`/auth?${params.toString()}`);
    };

    return(
        <section className="min-h-screen grid place-items-center p-4 bg-white">
            <div className="w-full max-w-2xl rounded-xl border border-blue-300 bg-[#f8fbff] p-5 sm:p-6">
                <h1 className="m-0 text-3xl font-bold tracking-tight text-black">QuickHand</h1>

                <div className="mt-4 flex flex-wrap gap-2">
                    <button
                        className={`rounded-full px-4 py-2 text-sm font-semibold border transition ${choice === "Signup" ? "border-blue-400 bg-blue-400 text-black" : "border-blue-300 bg-white text-black"}`}
                        type="button"
                        onClick={() => { handleChoiceChange("Signup"); }}
                        aria-pressed={choice === "Signup"}
                    >
                        SignUp
                    </button>

                    <button
                        className={`rounded-full px-4 py-2 text-sm font-semibold border transition ${choice === "Signin" ? "border-blue-400 bg-blue-400 text-black" : "border-blue-300 bg-white text-black"}`}
                        type="button"
                        onClick={() => { handleChoiceChange("Signin"); }}
                        aria-pressed={choice === "Signin"}
                    >
                        SignIn
                    </button>
                </div>

                <p className="mt-3 text-sm text-black">{choice === "Signin" ? "Welcome back" : "Create your account"}</p>

                <div className="mt-4">
                    {choice === "Signin" ? <SigninForm /> : <SignUpForm />}
                </div>
            </div>
        </section>
    )
}