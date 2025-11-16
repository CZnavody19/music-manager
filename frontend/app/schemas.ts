import { z } from "zod";

const dateTimeString = z.string().refine((date) => !isNaN(Date.parse(date)));

export const GeneralSettingsSchema = z.object({
    downloadPath: z.string().min(1, "Download path is required."),
    tempPath: z.string().min(1, "Temporary path is required."),
});

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

export const TidalSettingsSchema = z.object({
    authTokenType: z.string().min(1, "Auth token type is required."),
    authAccessToken: z.string().min(1, "Auth access token is required."),
    authRefreshToken: z.string().min(1, "Auth refresh token is required."),
    authExpiresAt: dateTimeString,
    authClientID: z.string().min(1, "Auth client ID is required."),
    authClientSecret: z.string().min(1, "Auth client secret is required."),
    downloadTimeout: z.number().min(1, "Download timeout must be at least 1 second."),
    downloadRetries: z.number().min(1, "Download retries must be at least 1."),
    downloadThreads: z.number().min(1, "Download threads must be at least 1."),
    audioQuality: z.string().min(1, "Audio quality is required.").refine((val) =>
        ["LOW", "HIGH", "LOSSLESS", "HI_RES_LOSSLESS"].includes(val),
        "Audio quality must be one of: LOW, HIGH, LOSSLESS, HI_RES_LOSSLESS."
    ),
});

export const YouTubeMatchSchema = z.object({
    trackId: z.uuidv4("Recording ID must be a valid UUID."),
});