package tdealerService

import (
	"net/http"

	"bitbucket.org/tekion/tbaas/apiContext"
	"bitbucket.org/tekion/tbaas/tapi"
	"bitbucket.org/tekion/tdealer/dealer"
	"github.com/gorilla/context"
	"github.com/unrolled/render"
)

//GetDealerByID :  Get Dealer By Id
func GetDealerByID(w http.ResponseWriter, r *http.Request) {

	render := render.New()

	// Get context
	ctx := context.Get(r, "apiContext").(apiContext.APIContext)

	dealerResponse, err := dealer.GetDealerByID(ctx, ctx.DealerID)
	if err != nil {
		res := tapi.ConstructResponse(http.StatusNotFound, err.Error(), nil)
		render.JSON(w, http.StatusNotFound, res)
		return
	}
	res := tapi.ConstructResponse(http.StatusOK, "Dealer Response ", dealerResponse)
	render.JSON(w, http.StatusOK, res)

}
