package config

// Config struct fills the parameters of request or user input
type Config struct {
	Endpoint string // CSI endpoint
	NodeID   string // CSI node ID
}

// NewConfig returns config struct to initialize new driver
func NewConfig() *Config {
	return &Config{}
}
