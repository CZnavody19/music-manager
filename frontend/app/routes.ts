import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
    index("routes/dashboard.tsx"),
    route("sources", "routes/sources.tsx", [
        route("match/:videoId", "routes/matching/youtube.tsx"),
    ]),
    route("integrations", "routes/integrations.tsx"),
    route("services", "routes/services.tsx"),
    route("settings", "layouts/settings.tsx", [
        index("routes/settings/general.tsx"),
        route("youtube", "routes/settings/youtube.tsx"),
        route("discord", "routes/settings/discord.tsx"),
        route("plex", "routes/settings/plex.tsx"),
    ]),

    route("login", "routes/auth/login.tsx"),
    route("logout", "routes/auth/logout.ts"),
] satisfies RouteConfig;
