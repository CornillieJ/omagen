package theme

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestColorsRoundTripsExtraAccents(t *testing.T) {
	dir := t.TempDir()
	p := Palette{
		Mode: "dark", Accent: "#ff2d95", Accent2: "#2de0c8", Accent3: "#f5d90a", Accent4: "#7a5cff", Accent5: "#22c55e",
		Selection: "#222222", Muted: "#333333",
		Background: "#444444", DarkBackground: "#555555", DarkerBackground: "#666666", LighterBackground: "#777777",
		Foreground: "#888888", DarkForeground: "#999999", LightForeground: "#aaaaaa", BrightForeground: "#bbbbbb",
		Red: "#cc0000", Yellow: "#ccaa00", Orange: "#cc6600", Green: "#00cc00", Cyan: "#00cccc", Blue: "#0000cc", Magenta: "#cc00cc", Brown: "#996633",
		BrightRed: "#ff0000", BrightYellow: "#ffff00", BrightGreen: "#00ff00", BrightCyan: "#00ffff", BrightBlue: "#0000ff", BrightMagenta: "#ff00ff",
	}
	if err := WriteColors(dir, p); err != nil {
		t.Fatal(err)
	}
	got, err := ReadColors(dir)
	if err != nil {
		t.Fatal(err)
	}
	if got.Accent2 != "#2de0c8" || got.Accent3 != "#f5d90a" || got.Accent4 != "#7a5cff" || got.Accent5 != "#22c55e" {
		t.Fatalf("extra accents = %+v, want Accent2=#2de0c8 Accent3=#f5d90a Accent4=#7a5cff Accent5=#22c55e", got)
	}
}

func TestColorsOmitAccent2WhenUnset(t *testing.T) {
	dir := t.TempDir()
	p := Palette{
		Mode: "dark", Accent: "#ff2d95", Selection: "#222222", Muted: "#333333",
		Background: "#444444", DarkBackground: "#555555", DarkerBackground: "#666666", LighterBackground: "#777777",
		Foreground: "#888888", DarkForeground: "#999999", LightForeground: "#aaaaaa", BrightForeground: "#bbbbbb",
		Red: "#cc0000", Yellow: "#ccaa00", Orange: "#cc6600", Green: "#00cc00", Cyan: "#00cccc", Blue: "#0000cc", Magenta: "#cc00cc", Brown: "#996633",
		BrightRed: "#ff0000", BrightYellow: "#ffff00", BrightGreen: "#00ff00", BrightCyan: "#00ffff", BrightBlue: "#0000ff", BrightMagenta: "#ff00ff",
	}
	if err := WriteColors(dir, p); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "colors.toml"))
	if err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"accent2", "accent3", "accent4", "accent5"} {
		if strings.Contains(string(data), key) {
			t.Fatalf("colors.toml unexpectedly contains %s:\n%s", key, data)
		}
	}
	got, err := ReadColors(dir)
	if err != nil {
		t.Fatal(err)
	}
	if got.Accent2 != "" || got.Accent3 != "" || got.Accent4 != "" || got.Accent5 != "" {
		t.Fatalf("extra accents = %+v, want all empty", got)
	}
}

func TestReadSharedWindowOpacityRecoversMatchingHyprlandValues(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "hyprland.lua"), []byte(`
hl.config({
  decoration = {
    active_opacity = 0.72,
    inactive_opacity = 0.72,
  },
})
`), 0o644); err != nil {
		t.Fatal(err)
	}

	opacity, err := ReadSharedWindowOpacity(dir)
	if err != nil {
		t.Fatal(err)
	}
	if opacity == nil || *opacity != 72 {
		t.Fatalf("window opacity = %#v, want 72", opacity)
	}
}

func TestReadSharedWindowOpacityLeavesAsymmetricOrMissingValuesUntouched(t *testing.T) {
	for name, source := range map[string]string{
		"asymmetric": "active_opacity = 0.72,\\ninactive_opacity = 0.60,\\n",
		"missing":    "active_opacity = 0.72,\\n",
	} {
		t.Run(name, func(t *testing.T) {
			dir := t.TempDir()
			if err := os.WriteFile(filepath.Join(dir, "hyprland.lua"), []byte(source), 0o644); err != nil {
				t.Fatal(err)
			}
			opacity, err := ReadSharedWindowOpacity(dir)
			if err != nil {
				t.Fatal(err)
			}
			if opacity != nil {
				t.Fatalf("window opacity = %d, want no shared value", *opacity)
			}
		})
	}
}
