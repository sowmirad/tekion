package assets

import (
	"encoding/json"
	"net/http"

	"fmt"

	"bitbucket.org/tekion/erratum"
	"bitbucket.org/tekion/tbaas/apiContext"
	mMgr "bitbucket.org/tekion/tbaas/mongoManager"
	"bitbucket.org/tekion/tbaas/tapi"
	"github.com/pkg/errors"
)

var serviceID = "assets"
var assetCol = "Asset"

//func context(r *http.Request) *apiContext.TContext {
//	ctx := new(apiContext.TContext)
//
//	ctx.Tenant = r.Header.Get("tenantName")
//	ctx.RequestID = r.Header.Get("requestID")
//	if len(ctx.RequestID) == 0 {
//		ctx.RequestID = uuid.NewUUID()
//	}
//	ctx.CorrelationID = r.Header.Get("correlationID")
//	if len(ctx.CorrelationID) == 0 {
//		ctx.CorrelationID = uuid.NewUUID()
//	}
//	return ctx
//}

func assetsH(w http.ResponseWriter, r *http.Request) {
	ctx := apiContext.UpgradeCtx(r.Context())

	arb := new(assetsReqBody)
	if err := json.NewDecoder(r.Body).Decode(&arb); err != nil {
		err = errors.Wrap(err, " failed to decode '/assets' payload ")
		tapi.HTTPErrorResponse(ctx, w, serviceID, erratum.ErrorDecodingPayload, err)
		return
	}

	findQ, err := arb.findQ()
	if err != nil {
		err = errors.Wrap(err, " failed to create find query to search assets ")
		tapi.HTTPErrorResponse(ctx, w, serviceID, erratum.ErrorDecodingPayload, err)
		return
	}
	searchQ := arb.searchQ()

	assets := make([]assets, 0, 0)
	if err = mMgr.ReadAll(ctx.Tenant, assetCol, findQ, searchQ, &assets); err != nil {
		err = errors.Wrap(err, " failed to read assets from db ")
		tapi.HTTPErrorResponse(ctx, w, serviceID, erratum.ErrorQueryingDB, err)
		return
	}
	if len(assets) == 0 {
		err = fmt.Errorf(" no assets found in db, findQ:%+v, request:%+v ", findQ, arb)
		tapi.HTTPErrorResponse(ctx, w, serviceID, erratum.ErrorQueryingDB, err)
		return
	}

	if len(assets) > 1 {
		err = fmt.Errorf(" multiple assets returned from db, findQ:%+v, request:%+v ", findQ, arb)
		tapi.HTTPErrorResponse(ctx, w, serviceID, erratum.ErrorQueryingDB, err)
		return
	}

	tapi.HTTPResponse(ctx, w, http.StatusOK, "assets", assets[0])
}
