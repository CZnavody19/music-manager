package domain

type DiscordConfig struct {
	Enabled    bool
	WebhookURL string
}

type DiscordMessageField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type DiscordMessage struct {
	Title       string                `json:"title"`
	Description *string               `json:"description"`
	Color       *int                  `json:"color"`
	Fields      []DiscordMessageField `json:"fields"`
}

type DiscordMessageSchema struct {
	Embeds []DiscordMessage `json:"embeds"`
}
