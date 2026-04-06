import Link from "next/link";
import { footerProps } from "@/types/props";

function FooterTemplate (props: footerProps) {
    return(
        <section className="text-white py-4 px-4 mt-auto w-full flex flex-wrap items-center justify-between gap-6 text-center md:text-left">
            <div className="flex flex-col items-center md:items-start gap-1 flex-1 min-w-0">
                <h1 className="text-lg font-bold">QuickHand</h1>
                <p className="text-xs text-gray-300">Hire handymen in the quickest fashion</p> 
                <p className="text-xs text-gray-400">copyright@QuickHand2025</p>   
            </div>
            <div className="flex flex-col items-center gap-1 flex-1 min-w-0">
                <h3 className="text-sm font-semibold">Contact</h3>
                <h5 className="text-xs text-gray-300">quickhand2025@gmail.com</h5>
                <h5 className="text-xs text-gray-300">+91 1234567890</h5>
            </div>
            <div className="flex flex-wrap justify-center md:justify-end gap-2 flex-1 min-w-0">
                {
                    props.buttons.map((btn, id) => (
                        btn.onClick
                            ? <button key={id} onClick={btn.onClick} className="px-1 text-white text-xs rounded transition-all duration-200 hover:scale-105 hover:shadow-lg">
                                {btn.buttonName}
                            </button>
                            : <Link href={btn.url!} key={id}>
                                <button className="px-1 text-white text-xs rounded transition-all duration-200 hover:scale-105 hover:shadow-lg">
                                    {btn.buttonName}
                                </button>
                            </Link>
                    ))
                }
            </div>
        </section>
    )
}

export default FooterTemplate;