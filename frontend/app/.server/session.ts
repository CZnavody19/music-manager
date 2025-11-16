import { createCookieSessionStorage, type Session } from "react-router";

export const sessionStorage = createCookieSessionStorage({
    cookie: {
        httpOnly: true,
        name: "session",
        path: "/",
        sameSite: "lax",
        secrets: [process.env.SESSION_SECRET!],
        // secure: process.env.NODE_ENV === "production",
    },
});

export const { commitSession, destroySession, getSession: _getSession } = sessionStorage;

export async function getSession(request: Request) {
    return await _getSession(request.headers.get("Cookie"));
}

export async function saveSession(session: Session) {
    return {
        headers: {
            "Set-Cookie": await commitSession(session),
        },
    }
}