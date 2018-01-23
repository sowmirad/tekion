package dealerService

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/tekion/tbaas/consulhelper"
	"bitbucket.org/tekion/tbaas/log"
	com "bitbucket.org/tekion/tenums/common"
	l "bitbucket.org/tekion/tenums/login"
)

func userByUserName(ctx *customCtx) (*user, error) {
	endpoint := consulhelper.GetServiceNodes(l.ServiceID) + l.UserByUserName + ctx.UserName

	res, err := callAPI(ctx, com.AppJSON, http.MethodGet, endpoint, nil, "")
	if err != nil {
		err = fmt.Errorf("call to %s failed, error: %v", endpoint, err)
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return nil, err
	}
	defer res.Body.Close()
	var userByUserName userByUserNameRes
	// reading from response body until an error or EOF
	err = json.NewDecoder(res.Body).Decode(&userByUserName)
	if err != nil {
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return nil, err
	}

	return &userByUserName.Data, err
}

func userByID(ctx *customCtx, userID string) (*user, error) {
	endpoint := consulhelper.GetServiceNodes(l.ServiceID) + l.UserByID + userID

	res, err := callAPI(ctx, com.AppJSON, http.MethodGet, endpoint, nil, "")
	if err != nil {
		err = fmt.Errorf("call to %s failed, error: %v", endpoint, err)
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return nil, err
	}
	defer res.Body.Close()
	var userByID userByIDRes
	// reading from response body until an error or EOF
	err = json.NewDecoder(res.Body).Decode(&userByID)
	if err != nil {
		log.GenericError(ctx.Tenant, ctx.DealerID, ctx.UserName, err)
		return nil, err
	}

	return &userByID.Data, err
}
