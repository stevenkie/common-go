package config

type (
	// NewConfigOptions options for initializing new config instance
	NewConfigOptions struct {
		// DefaultName used for default config name that will be used
		DefaultName string

		// Defaults used for default value for each config
		Defaults map[string]interface{}
		IsWatch  bool
	}
)
