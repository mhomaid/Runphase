package version

import "testing"

func TestModulePath(t *testing.T) {
	const want = "github.com/mhomaid/runphase"
	if ModulePath != want {
		t.Fatalf("ModulePath = %q, want %q", ModulePath, want)
	}
}
