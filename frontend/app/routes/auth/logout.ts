import { getSession, saveSession } from "~/.server/session";
import type { Route } from "./+types/logout";
import { redirect } from "react-router";

export async function loader({ request }: Route.LoaderArgs) {
    const redir = new URL(request.url).searchParams.get("redirect") ?? "/";
    const session = await getSession(request);

    if (!session.has("token")) {
        return redirect(redir);
    }

    session.unset("token");

    return redirect(redir, await saveSession(session));
}