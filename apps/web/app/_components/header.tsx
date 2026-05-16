"use client"


import { authClient, useSession } from "@/packages/auth/auth-client"
import { useRouter } from "next/navigation";

export default function LandingHeader(){
    const { data: session, isPending } = useSession();
    const router = useRouter();

    return (
        <section className="bg-blue-400 px-5 py-5 font-semibold">
            <div className="flex justify-between px-5">
                <div>
                    <span className="text-3xl">
                        <span className="text-white">Quick</span>
                        <span className="text-black">Hand</span>
                    </span>
                </div>
                <div>
                    <nav className="flex justify-evenly gap-4 text-xl">
                        <button className="hover:px-3 hover:border-2  text-black hover:bg-white hover:border-black rounded-xl">Sign Up</button>
                        <button className="hover:px-3 hover:border-2 hover:bg-black hover:border-black rounded-xl">Sign In</button>
                        <button className="hover:px-3 hover:border-2 text-black hover:bg-white hover:border-black rounded-xl">About</button>
                    </nav>
                </div>
            </div>
        </section>
    )
}