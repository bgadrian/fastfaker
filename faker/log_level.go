package faker

import (
	"github.com/bgadrian/fastfaker/data"
)

// LogLevel will generate a random log level
// See data/LogLevels for list of available levels
func (f *Faker) LogLevel(logType string) string {
	if _, ok := data.LogLevels[logType]; ok {
		return f.getRandValue([]string{"log_level", logType})
	}

	return f.getRandValue([]string{"log_level", "general"})
}
