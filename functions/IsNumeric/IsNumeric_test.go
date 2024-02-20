package IsNumeric

import "testing"

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"Empty string", "", true},
		{"String with numbers", "12345", true},
		{"String with letters", "abc123", false},
		{"String with special characters", "!@#$%^&*()", false},
		{"String with mixed characters", "123abc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumeric(tt.s); got != tt.want {
				t.Errorf("isNumeric(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
