package generator

import "testing"

// TODO: add fuzzing to test edge cases
func TestShortCodeGenerator_Generate(t *testing.T) {
	tests := []struct {
		name     string
		g        *ShortCodeGenerator
		wantSize int
	}{
		{
			name:     "generate url code",
			g:        &ShortCodeGenerator{},
			wantSize: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &ShortCodeGenerator{}
			if got := g.Generate(); len(got) != tt.wantSize {
				t.Errorf("ShortCodeGenerator.Generate() = %v, want len %v", got, tt.wantSize)
			}
		})
	}
}
