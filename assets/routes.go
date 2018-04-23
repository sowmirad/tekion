package assets

import (
	"net/http"

	"bitbucket.org/tekion/tbaas/tapi"
)

func Start() {

	tapi.AddNoAuthRoutes(
		"assets",
		http.MethodPost,
		"/assets",
		assetsH,
	)
}
