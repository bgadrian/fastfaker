package faker

import (
	"fmt"
)

// AvatarURL will generate a random image URL of 80x80 px. Images Provided by pipsum.com
func (f *Faker) AvatarURL() string {
	return f.ImageURL(80, 80)
}

// ImageURL will generate a random Image Based Upon Height And Width. Images Provided by pipsum.com/
func (f *Faker) ImageURL(width int, height int) string {
	width = clampInt(width, 1, 2560)
	height = clampInt(height, 1, 1600)

	return fmt.Sprintf("http://pipsum.com/%dx%d.jpg", width, height)
}
