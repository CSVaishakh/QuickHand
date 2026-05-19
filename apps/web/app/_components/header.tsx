"use client"


import { useSession } from "@/packages/auth/auth-client"
import Link from "next/link";

export default function LandingHeader(){
    const { data: session } = useSession();

    return (
        <section className="bg-blue-400 px-5 py-5 font-semibold">
            <div className="flex justify-between px-5">
                <div>
                    <span className="text-3xl">
                        <span className="text-white">Quick</span>
                        <span className="text-black">Hand</span>
                    </span>
                </div>
                { session?
                    ( 
                        <div>
                            <nav className="flex justify-evenly gap-4 text-xl">
                                <Link href={"/auth/profile"}>
                                    <button className="hover:px-3 hover:border-2  text-black hover:bg-white hover:border-black rounded-xl">
                                        {session.user.name ?? "Profile"}
                                    </button>
                                </Link>
                                
                                <Link href={"/about"}>
                                    <button className="hover:px-3 hover:border-2 text-black hover:bg-white hover:border-black rounded-xl">
                                        About
                                    </button>
                                </Link>
                            </nav> 
                        </div>
                    ):(
                       <div>
                            <nav className="flex justify-evenly gap-4 text-xl">
                                <Link href={"/auth/sign-up"}>
                                    <button className="hover:px-3 hover:border-2  text-black hover:bg-white hover:border-black rounded-xl">
                                        Sign Up
                                    </button>
                                </Link>
                                
                                <Link href={"/auth/sign-in"}>
                                    <button className="hover:px-3 hover:border-2 hover:bg-black hover:border-black rounded-xl">
                                        Sign In
                                    </button>
                                </Link>
                                
                                <Link href={"/about"}>
                                    <button className="hover:px-3 hover:border-2 text-black hover:bg-white hover:border-black rounded-xl">
                                        About
                                    </button>
                                </Link>
                            </nav>
                        </div> 
                    )
                }
            </div>
        </section>
    )
}