package tdealerService

import (
	"net/http"

	"bitbucket.org/tekion/tbaas/tapi"
	"bitbucket.org/tekion/tbaas/apiContext"
	"github.com/gorilla/context"
	"github.com/unrolled/render"
	"bitbucket.org/tekion/tdealer/dealer"
)

//GetDealerById :  Get Dealer By Id
func GetDealerById(w http.ResponseWriter, r *http.Request) {

	render := render.New()

	// Get context
	ctx := context.Get(r, "apiContext").(apiContext.APIContext)

	dealerResponse, err := dealer.GetDealerById(ctx,ctx.DealerID)
	if err != nil {
		res := tapi.ConstructResponse(http.StatusNotFound, err.Error(), nil)
		render.JSON(w, http.StatusNotFound, res)
		return
	}
	res := tapi.ConstructResponse(http.StatusOK, "Dealer Response ", dealerResponse)
	render.JSON(w, http.StatusOK, res)

}


