import Link from "next/link";
import { headerProps } from "@/types/props";

function headerTemplate( props: headerProps ) {
    return(
        <section className="bg-linear-to-r from-blue-400 to-blue-300 text-black px-10 py-7 flex justify-between items-stretch">
            <div className="font-medium text-2xl ">QuickHand</div>
            <nav className="flex gap-6">
                {
                    props.buttons.map((btn, id) => (
                        <Link href={btn.url} key={id}><button className="text-xl font-medium px-2 transition-all hover:bg-blue-400 hover:border-2 hover:border-blue-400 hover:rounded-lg">{btn.buttonName}</button></Link>
                    ))
                }
            </nav>
        </section>
    )   
}

export default headerTemplate;