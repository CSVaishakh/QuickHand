import "better-auth";

declare module "better-auth" {
  interface User {
    role?: "customer" | "electrician" | "admin";
  }

  interface Session {
    user: User;
  }
}