package theme

import (
	"reflect"
	"testing"
)

func TestPaletteActiveAccentsStopsAtFirstUnset(t *testing.T) {
	tests := []struct {
		name    string
		palette Palette
		want    []string
	}{
		{"accent only", Palette{Accent: "#111111"}, []string{"#111111"}},
		{
			"accent through accent5",
			Palette{Accent: "#111111", Accent2: "#222222", Accent3: "#333333", Accent4: "#444444", Accent5: "#555555"},
			[]string{"#111111", "#222222", "#333333", "#444444", "#555555"},
		},
		{
			"gap after accent2 ignores accent3+",
			Palette{Accent: "#111111", Accent2: "#222222", Accent4: "#444444"},
			[]string{"#111111", "#222222"},
		},
		{
			"invalid accent2 stops immediately",
			Palette{Accent: "#111111", Accent2: "not-a-color", Accent3: "#333333"},
			[]string{"#111111"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.palette.ActiveAccents()
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("ActiveAccents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaletteValidateAllowsUnsetExtraAccents(t *testing.T) {
	base := Palette{
		Mode: "dark", Accent: "#111111", Selection: "#222222", Muted: "#333333",
		Background: "#444444", DarkBackground: "#555555", DarkerBackground: "#666666", LighterBackground: "#777777",
		Foreground: "#888888", DarkForeground: "#999999", LightForeground: "#AAAAAA", BrightForeground: "#BBBBBB",
		Red: "#CC0000", Yellow: "#CCAA00", Orange: "#CC6600", Green: "#00CC00", Cyan: "#00CCCC", Blue: "#0000CC", Magenta: "#CC00CC", Brown: "#996633",
		BrightRed: "#FF0000", BrightYellow: "#FFFF00", BrightGreen: "#00FF00", BrightCyan: "#00FFFF", BrightBlue: "#0000FF", BrightMagenta: "#FF00FF",
	}
	if err := base.Validate(); err != nil {
		t.Fatalf("Validate() with no extra accents set = %v, want nil", err)
	}

	base.Accent3 = "not-a-color"
	if err := base.Validate(); err == nil {
		t.Fatal("Validate() accepted an invalid Accent3")
	}
}
