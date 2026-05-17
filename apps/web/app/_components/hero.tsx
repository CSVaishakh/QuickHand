"use client"

import Image from "next/image";
import Link from "next/link";

export default function LandingHero () {
    return(
        <section className="min-h-screen bg-blue-400">
            <div className="flex flex-col items-center justify-center p-20">    
                <div className="inline-flex items-center self-center gap-3 px-4 py-2 border-2 border-black rounded-full bg-white text-2xl text-black font-semibold">
                    <span className="flex"><Image src="/assets/pin.svg" alt={"📍"} width={25} height={25}/>Local</span>
                    <span className="flex"><Image src="/assets/tick.svg" alt={""} width={25} height={25}/>Trusted</span>
                    <span className="flex"><Image src="/assets/lightning.svg" alt={"📍"} width={25} height={25}/>Fast</span>
                </div>
                <div className="flex flex-col items-center text-8xl text-center font-bold gap-5 p-10">
                    <span>Hire <span className="text-black">Handymen</span> in the</span>
                    <span>quickest <span className="text-black">fashion.</span></span>
                </div>
                
                <div className="flex flex-col w-fit items-center gap-3 text-3xl font-bold">
                    <Link href={"/dashboard"}>
                        <button className="w-fit border-2 rounded-xl px-5 hover:bg-white hover:text-blue-400">
                            Find <span className="text-black">Handymen</span>
                        </button>
                    </Link>
                    <Link href={"/dashboard"}>
                        <button className="w-fit border-2 rounded-xl px-5 hover:bg-white hover:text-blue-400">
                            Find <span className="text-black">Work</span>
                        </button>
                    </Link>    
                </div>
            </div>
        </section>
    )
}