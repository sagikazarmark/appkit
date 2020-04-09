package buildinfo

import (
	"runtime"
	"testing"
)

func TestNew(t *testing.T) {
	buildinfo := New("version", "commit", "date")

	if want, have := "version", buildinfo.Version; want != have {
		t.Errorf("unexpected version\nexpected: %s\nactual:   %s", want, have)
	}

	if want, have := "commit", buildinfo.CommitHash; want != have {
		t.Errorf("unexpected commit hash\nexpected: %s\nactual:   %s", want, have)
	}

	if want, have := "date", buildinfo.BuildDate; want != have {
		t.Errorf("unexpected build date\nexpected: %s\nactual:   %s", want, have)
	}
}

func TestBuildInfo_Fields(t *testing.T) {
	buildinfo := New("version", "commit", "date")

	actualFields := buildinfo.Fields()

	expectedFields := map[string]interface{}{
		"version":     "version",
		"commit_hash": "commit",
		"build_date":  "date",
		"go_version":  runtime.Version(),
		"os":          runtime.GOOS,
		"arch":        runtime.GOARCH,
		"compiler":    runtime.Compiler,
	}

	if want, have := len(expectedFields), len(actualFields); want != have {
		t.Errorf("unexpected fields\nexpected: %s\nactual:   %s", expectedFields, actualFields)
	}

	for key, value := range expectedFields {
		if want, have := value, actualFields[key]; want != have {
			t.Errorf("unexpected field value for %q\nexpected: %s\nactual:   %s", key, want, have)
		}
	}
}
