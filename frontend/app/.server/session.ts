import { createCookieSessionStorage } from "react-router";

export const sessionStorage = createCookieSessionStorage({
    cookie: {
        httpOnly: true,
        name: "session",
        path: "/",
        sameSite: "lax",
        secrets: [process.env.SESSION_SECRET!],
        secure: process.env.NODE_ENV === "production",
    },
});

export const { commitSession, destroySession, getSession } = sessionStorage;