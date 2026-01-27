package website

// FontConfig holds Google Fonts configuration for a font family.
type FontConfig struct {
	URL    string
	Family string
}

// fontRegistry maps FontFamily constants to their configurations.
var fontRegistry = map[FontFamily]FontConfig{
	FontGoogleSansFlex: {
		URL:    "https://fonts.googleapis.com/css2?family=Google+Sans+Flex:opsz,wght@6..144,1..1000&family=Inter:ital,opsz,wght@0,14..32,100..900;1,14..32,100..900&display=swap",
		Family: "'Google Sans Flex', 'Inter', sans-serif",
	},
}

// GetFontConfig returns the font configuration for a given FontFamily.
// Returns GoogleSansFlex configuration as fallback for unknown fonts.
func GetFontConfig(font FontFamily) FontConfig {
	if cfg, ok := fontRegistry[font]; ok {
		return cfg
	}
	return fontRegistry[FontGoogleSansFlex]
}
