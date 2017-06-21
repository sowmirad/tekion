package tdealerService

import (
	"net/http"

	"bitbucket.org/tekion/erratum"
	"bitbucket.org/tekion/tbaas/apiContext"
	"bitbucket.org/tekion/tbaas/tapi"
	"bitbucket.org/tekion/tdealer/dealer"
	"github.com/gorilla/context"
)

var (
	serviceName = "estimate"
)

//GetDealerByID :  Get Dealer By Id
func GetDealerByID(w http.ResponseWriter, r *http.Request) {
	// Get context
	ctx := context.Get(r, "apiContext").(apiContext.APIContext)

	dealerResponse, err := dealer.GetDealerByID(ctx, ctx.DealerID)
	if err != nil {
		tapi.WriteHTTPErrorResponse(w, serviceName, erratum.ErrorQueryingDB, err)
		return
	}

	tapi.WriteHTTPResponse(w, http.StatusOK, "Dealer Response ", dealerResponse)
	return
}

func getDamageTypes(w http.ResponseWriter, r *http.Request) {
	ctx := context.Get(r, "apiContext").(apiContext.APIContext)
	res, err := dealer.GetDamageTypes(ctx, ctx.DealerID)
	if err != nil {
		tapi.WriteHTTPErrorResponse(w, serviceName, erratum.ErrorQueryingDB, err)
		return
	}

	tapi.WriteHTTPResponse(w, http.StatusOK, "GetDamageTypes Response ", res)
	return
}
