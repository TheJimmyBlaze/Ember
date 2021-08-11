package api

import (
	"fmt"
	"net/http"

	"github.com/thejimmyblaze/ember/version"
)

func (h *APIHandler) Version(w http.ResponseWriter, r *http.Request) {
	info := h.buildVersionInfo()
	fmt.Fprint(w, info)
}

func (h *APIHandler) buildVersionInfo() string {

	info := fmt.Sprintf("Ember - X.509 Crypto Service - %s", version.BuildVersion)
	info = fmt.Sprintf("%s\nDistributed under the MIT licence: github.com/thejimmyblaze/ember", info)
	info = fmt.Sprintf("%s\n", info)
	info = fmt.Sprintf("%s\n\nVersion Information:", info)
	info = fmt.Sprintf("%s\nBuild Version: %s", info, version.BuildVersion)
	info = fmt.Sprintf("%s\nBuild Time: %s", info, version.BuildTime)
	info = fmt.Sprintf("%s\nBuild Hash: %s", info, version.BuildHash)

	return info
}
