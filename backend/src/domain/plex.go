package domain

type PlexConfig struct {
	Enabled   bool
	Protocol  string
	Host      string
	Port      int64
	Token     string
	LibraryID int64
}
