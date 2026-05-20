"use client";

import Link from "next/link";
import Image from "next/image";
import { authClient } from "@/packages/auth/auth-client";

type UserDropdownProps = {
  session: {
    user: {
      name?: string | null;
      image?: string | null;
    };
  };
};

export function UserDropdown({ session }: UserDropdownProps) {
  return (
    <details className="relative text-black">
      <summary className="flex items-center gap-2 cursor-pointer list-none [&::-webkit-details-marker]:hidden">
        <Image
          src={session.user.image || "/assets/avatar.svg"}
          alt="user"
          width={22}
          height={22}
          className="rounded-full"
        />

        <span>{session.user.name}</span>
      </summary>

      <div className="absolute right-0 mt-2 w-40 overflow-hidden rounded-xl border border-gray-300 bg-white shadow-md">
        <Link
          href="/auth/profile"
          className="block px-4 py-3 hover:bg-gray-100"
        >
          Profile
        </Link>

        <button
          onClick={() => authClient.signOut()}
          className="w-full px-4 py-3 text-left hover:bg-gray-100"
        >
          Sign Out
        </button>
      </div>
    </details>
  );
}