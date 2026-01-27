package website

import "testing"

func TestGetFontConfig(t *testing.T) {
	tests := []struct {
		name string
		font FontFamily
		want string // Family name
	}{
		{
			name: "existing font",
			font: FontGoogleSansFlex,
			want: "'Google Sans Flex', 'Inter', sans-serif",
		},
		{
			name: "unknown font fallback",
			font: "Comic Sans",
			want: "'Google Sans Flex', 'Inter', sans-serif",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetFontConfig(tt.font)
			if got.Family != tt.want {
				t.Errorf("GetFontConfig().Family = %v, want %v", got.Family, tt.want)
			}
		})
	}
}
