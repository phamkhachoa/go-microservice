package initialize

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go-ecommerce-backend-api/global"
	"os"
	"regexp"
)

// LoadConfig loads configuration with environment variable expansion
func LoadConfig(configPath string, envPath string) {
	// Load .env file if specified
	if envPath != "" {
		err := godotenv.Load(envPath)
		if err != nil {
			// Just log warning but continue, as env vars might be set elsewhere
			fmt.Printf("Warning: Error loading .env file: %v\n", err)
		}
	}

	// Initialize Viper
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")

	// Read the config file
	if err := v.ReadInConfig(); err != nil {
		fmt.Errorf("failed to read config file: %w", err)
	}

	// Process all config values to replace environment variables
	expandEnvVariables(v)

	// Create a new config instance
	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Errorf("failed to read config file: %w", err)
	}
}

// expandEnvVariables processes config values and replaces ${VAR:default} patterns with environment values
func expandEnvVariables(v *viper.Viper) {
	// Get all settings
	settings := v.AllSettings()

	// Process settings recursively
	processMap(v, settings, "")
}

// processMap processes a map of config values
func processMap(v *viper.Viper, inputMap map[string]interface{}, prefix string) {
	for key, value := range inputMap {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "." + key
		}

		// If value is a map, process it recursively
		if nestedMap, ok := value.(map[string]interface{}); ok {
			processMap(v, nestedMap, fullKey)
		} else if strValue, ok := value.(string); ok {
			// If it's a string, check for ${VAR:default} pattern
			expandedValue := expandEnvValue(strValue)
			if expandedValue != strValue {
				v.Set(fullKey, expandedValue)
			}
		}
	}
}

// expandEnvValue expands ${VAR:default} in a string with environment variable values
func expandEnvValue(value string) string {
	// Regex to match ${VAR:default} pattern
	re := regexp.MustCompile(`\${([^:}]+)(?::([^}]*))?}`)

	return re.ReplaceAllStringFunc(value, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 3 {
			return match
		}

		envVarName := submatches[1]
		defaultValue := submatches[2]

		// Get environment variable or use default
		envValue := os.Getenv(envVarName)
		if envValue != "" {
			return envValue
		}
		return defaultValue
	})
}
