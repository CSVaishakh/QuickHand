"use client"

import type { customer } from "@/types/types";
import { useState } from "react";
import { useRouter } from "next/navigation";

export default function CustomerForm () {
    const router = useRouter();

    const [form, setForm] = useState<customer>({
        name: "",
        email: "",
        password: ""
    });

    const [loading, setLoading] = useState(false);
    const [error, setError] = useState("");
    const [success, setSuccess] = useState(false);

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setError("");
        setSuccess(false);
        setLoading(true);

        try {
            const baseUrl = (
                process.env.NEXT_PUBLIC_API_URL ??
                process.env.NEXT_PUBLIC_BETTER_AUTH_URL ??
                (typeof window !== "undefined" ? window.location.origin : "")
            ).replace(/\/$/, "");

            const response = await fetch(`${baseUrl}/customer/sign-up`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                credentials: "include",
                body: JSON.stringify(form),
            });

            if (!response.ok) {
                const contentType = response.headers.get("content-type") ?? "";
                let message = "Signup failed";

                if (contentType.includes("application/json")) {
                    const data = await response.json();
                    message = data.error ?? data.message ?? message;
                }

                throw new Error(message);
            }

            setSuccess(true);
            setForm({
                name: "",
                email: "",
                password: "",
            });
            router.push("/dashboard");
        } catch (err) {
            setError(err instanceof Error ? err.message : "Something went wrong");
        } finally {
            setLoading(false);
        }
    };

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
            const {name, value} = e.target;
            setForm((prev) => ({
                ...prev,
                [name]: value
            }));
    }

    return(
        <form onSubmit={handleSubmit} className="grid gap-4">
            <div className="grid gap-3">
                <div className="grid gap-1.5">
                    <h3 className="m-0 text-sm font-semibold text-black">Name</h3>
                    <input  type="text"
                            name="name"
                            value={form.name}
                            placeholder="Full Name"
                            onChange={handleChange}
                            className="w-full rounded-lg border border-blue-300 bg-white px-3 py-2.5 text-sm text-black outline-none transition focus:border-blue-400"
                    />
                </div>

                <div className="grid gap-1.5">
                    <h3 className="m-0 text-sm font-semibold text-black">Email</h3>
                    <input  type="email"
                            name="email" 
                            value={form.email} 
                            placeholder="example@gmail.com"
                            onChange={handleChange}
                            className="w-full rounded-lg border border-blue-300 bg-white px-3 py-2.5 text-sm text-black outline-none transition focus:border-blue-400"
                    />
                </div>

                <div className="grid gap-1.5">
                    <h3 className="m-0 text-sm font-semibold text-black">Password</h3>
                    <input  type="password" 
                            name="password"
                            value={form.password}
                            placeholder="Example@123"
                            onChange={handleChange}
                            className="w-full rounded-lg border border-blue-300 bg-white px-3 py-2.5 text-sm text-black outline-none transition focus:border-blue-400"
                    />
                </div>
            </div>
            <div className="grid gap-2">
                <button type="submit" disabled={loading} className="rounded-lg border border-blue-400 bg-blue-400 px-4 py-2.5 text-sm font-semibold text-black transition hover:bg-blue-300 disabled:cursor-not-allowed disabled:opacity-65">
                    {loading ? "Creating account..." : "Create Account"}
                </button>
                {error ? <p className="m-0 text-sm text-black">{error}</p> : null}
                {success ? <p className="m-0 text-sm text-black">Account created successfully.</p> : null}
            </div>
        </form>
    )
}