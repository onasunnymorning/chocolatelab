package batch

import "testing"

func TestNewOrigin(t *testing.T) {
	name := "Hello\nWorld"
	expectedName := "Hello World"

	origin := NewOrigin(name)

	if origin.Name != expectedName {
		t.Errorf("NewOrigin(%q) = %q; expected %q", name, origin.Name, expectedName)
	}
}
