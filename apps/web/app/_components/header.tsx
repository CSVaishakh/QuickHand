"use client"

import HeaderTemplate from "@/components/header.template";
import { headerProps } from "@/types/props";
import { authClient, useSession } from "@/packages/auth/auth-client"
import { useMemo } from "react";
import { useRouter } from "next/navigation";

export default function LandingHeader(){
    const { data: session, isPending } = useSession();
    const router = useRouter();

    const props = useMemo<headerProps>(() => {
        if (isPending) {
            return {
                buttons: [{ buttonName: "About", url: "/about" }],
            };
        }

        if (!session) {
            return {
                buttons: [
                    { buttonName: "About", url: "/about" },
                    { buttonName: "Sign Up", url: "/auth?mode=signup" },
                    { buttonName: "Sign In", url: "/auth?mode=signin" },
                ],
            };
        }

        return {
            buttons: [
                { buttonName: "About", url: "/about" },
                { buttonName: "Dashboard", url: "/dashboard" },
                {
                    buttonName: "Sign Out",
                    onClick: async () => {
                        await authClient.signOut();
                        router.push("/auth");
                    },
                },
            ],
        };
    }, [session, isPending, router]);

    return <HeaderTemplate {...props}/>;
}