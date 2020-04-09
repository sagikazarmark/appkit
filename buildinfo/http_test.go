package buildinfo

import (
	"encoding/json"
	"net/http/httptest"
	"runtime"
	"testing"
)

func TestHTTPHandler(t *testing.T) {
	buildinfo := New("version", "commit", "date")

	server := httptest.NewServer(HTTPHandler(buildinfo))
	defer server.Close()

	resp, err := server.Client().Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var actualFields map[string]interface{}

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&actualFields)
	if err != nil {
		t.Fatal(err)
	}

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
