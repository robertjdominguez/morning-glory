package types

// Each value corresponds to an env var.
type Config struct {
	TodoistApiToken  string
	TodoistProjectId string
	GoogleConfigFile string
	OpenAiApiKey     string
	Prompt           string
}
