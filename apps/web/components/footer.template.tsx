import Link from "next/link";
import { footerProps } from "@/types/props";

function footerTemplate (props: footerProps) {
    return(
        <section>
            <div>
                <div>
                    <h1>QuickHand</h1>
                    <p>Hire handymen in the quickest fashion</p> 
                    <p>copyright@QuickHand2025</p>   
                </div>
                <div>
                    <h3>Contact</h3>
                    <h5>quickhand2025@gmail.com</h5>
                    <h5>+91 1234567890</h5>
                </div>
            </div>            
            {
                props.buttons.map((btn, id) => (
                    <Link href={btn.url} key={id}><button>{btn.buttonName}</button></Link>
                ))
            }
        </section>
    )
}

export default footerTemplate;