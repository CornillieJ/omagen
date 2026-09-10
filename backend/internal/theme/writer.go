package theme

import (
	"fmt"
	"path/filepath"

	"github.com/prettyletto/omagen/backend/internal/fsutil"
)

func WriteColors(
	themeDir string,
	palette Palette,
) error {
	if err := palette.Validate(); err != nil {
		return fmt.Errorf(
			"validate palette: %w",
			err,
		)
	}

	path := filepath.Join(
		themeDir,
		"colors.toml",
	)

	var extraAccentLines string
	for _, extra := range []struct {
		key   string
		value string
	}{
		{"accent2", palette.Accent2},
		{"accent3", palette.Accent3},
		{"accent4", palette.Accent4},
		{"accent5", palette.Accent5},
	} {
		if extra.value != "" {
			extraAccentLines += fmt.Sprintf("%s = %q\n", extra.key, extra.value)
		}
	}

	content := fmt.Sprintf(
		`mode = %q

accent = %q
selection = %q
muted = %q

background = %q
dark_background = %q
darker_background = %q
lighter_background = %q

foreground = %q
dark_foreground = %q
light_foreground = %q
bright_foreground = %q

red = %q
yellow = %q
orange = %q
green = %q
cyan = %q
blue = %q
magenta = %q
brown = %q

bright_red = %q
bright_yellow = %q
bright_green = %q
bright_cyan = %q
bright_blue = %q
bright_magenta = %q
`,
		palette.Mode,

		palette.Accent,
		palette.Selection,
		palette.Muted,

		palette.Background,
		palette.DarkBackground,
		palette.DarkerBackground,
		palette.LighterBackground,

		palette.Foreground,
		palette.DarkForeground,
		palette.LightForeground,
		palette.BrightForeground,

		palette.Red,
		palette.Yellow,
		palette.Orange,
		palette.Green,
		palette.Cyan,
		palette.Blue,
		palette.Magenta,
		palette.Brown,

		palette.BrightRed,
		palette.BrightYellow,
		palette.BrightGreen,
		palette.BrightCyan,
		palette.BrightBlue,
		palette.BrightMagenta,
	)
	content += extraAccentLines

	if err := fsutil.AtomicWriteFile(path, []byte(content), 0o644); err != nil {
		return fmt.Errorf(
			"write colors.toml: %w",
			err,
		)
	}

	return nil
}
