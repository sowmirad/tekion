package assets

import (
	"net/http"

	"bitbucket.org/tekion/tbaas/tapi"
)

func Start() {

	tapi.AddNoAuthServiceRoute(
		"assets",
		http.MethodPost,
		"/assets",
		assetsH,
	)
}
