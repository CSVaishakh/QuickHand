import headerTemplate from "@/components/header.template";
import { headerProps } from "@/types/props";

const props: headerProps = { buttons: [{ buttonName: "About", url: "/about" },{buttonName: "SignUp", url: "/auth"},{buttonName: "SignIn",url:"/auth"}] }

export default function LandingHeader(){
    return(headerTemplate(props))
}