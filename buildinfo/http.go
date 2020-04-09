package buildinfo

import (
	"encoding/json"
	"net/http"
)

// HTTPHandler returns an HTTP handler for version information.
func HTTPHandler(buildInfo BuildInfo) http.Handler {
	var body []byte

	var err error

	body, err = json.Marshal(buildInfo)
	if err != nil {
		panic(err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(body)
	})
}
