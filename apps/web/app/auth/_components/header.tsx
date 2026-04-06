import HeaderTemplate from "@/components/header.template";
import { headerProps } from "@/types/props";

export default function AuthHeader() {
    const props: headerProps = {
        buttons: [{ buttonName: "Home", url: "/" }],
    };

    return <HeaderTemplate {...props} />;
}