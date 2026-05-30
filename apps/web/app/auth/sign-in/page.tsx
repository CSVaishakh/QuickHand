import Link from "next/link"
import Image from "next/image"

import SignInHeader from "./_components/header";
import SignInHero from "./_components/hero";

export default function SignIn() {
  return (
    <section className="min-h-screen flex flex-col items-center justify-center bg-blue-400">
        <main className="aspect-16/17 w-120 bg-white border-4 border-black rounded-2xl shadow-2xl flex flex-col justify-center">
            <SignInHeader/>
            <div className="px-4">
                <Link href={"/"}>
                    <button className="border-2 border-black rounded-xl px-6 py-2 text-black flex hover:bg-blue-200">
                        <Image src="/assets/left-arrow.svg" alt={"⬅️"} width={25} height={25}/>
                        back
                    </button>
                </Link>
            </div>
            <SignInHero />
            <div className="text-blue-400 flex justify-center py-2">
                <p>New to Quickhand? <Link href={"/auth/sign-up"}><button>Create Account!</button></Link></p>
                
            </div>
        </main>
    </section>
  );
}