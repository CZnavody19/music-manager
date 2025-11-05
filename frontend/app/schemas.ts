import { z } from "zod";

export const DiscordSettingsSchema = z.object({
    webhookURL: z.url("Please enter a valid URL for the webhook."),
});