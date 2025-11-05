import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
    index("routes/dashboard.tsx"),
    route("services", "routes/services.tsx"),
    route("settings", "layouts/settings.tsx", [
        index("routes/settings/general.tsx"),
        route("youtube", "routes/settings/youtube.tsx"),
        route("discord", "routes/settings/discord.tsx"),
        route("plex", "routes/settings/plex.tsx"),
    ]),
] satisfies RouteConfig;
