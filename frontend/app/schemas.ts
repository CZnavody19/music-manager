import { z } from "zod";

export const DiscordSettingsSchema = z.object({
    webhookURL: z.url("Please enter a valid URL for the webhook."),
});

export const PlexSettingsSchema = z.object({
    protocol: z.string("Protocol is required.").refine((val) => val === "http" || val === "https", "Protocol must be either 'http' or 'https'."),
    host: z.string().min(1, "Host is required."),
    port: z.number().min(1).max(65535).optional(),
    token: z.string().min(1, "Token is required."),
    libraryID: z.number("Library ID must be a number."),
});

export const AuthSchema = z.object({
    username: z.string().min(1, "Username is required."),
    password: z.string().min(8, "Password must be at least 8 characters long."),
});

export const YouTubeSettingsSchema = z.object({
    playlistID: z.string().min(1, "Playlist ID is required."),
});