package dealerService

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/tekion/tbaas/consulhelper"
	"bitbucket.org/tekion/tbaas/hwrap"
	log "bitbucket.org/tekion/tbaas/log/v1"
	com "bitbucket.org/tekion/tenums/common"
	l "bitbucket.org/tekion/tenums/login"
)

func userByUserName(ctx *customCtx) (*user, error) {
	endpoint := consulhelper.GetServiceNodes(ctx.TContext, l.ServiceID) + l.UserByUserName + ctx.UserName

	reqP := hwrap.RequestParams{
		Method:      http.MethodGet,
		ContentType: com.AppJSON,
		URL:         endpoint,
	}
	res, err := hwrap.HTTPSuccessRequest(ctx.TContext, reqP)
	if err != nil {
		err = fmt.Errorf("call to %s failed, error: %v", endpoint, err)
		log.GenericError(ctx.TContext, err, nil)
		return nil, err
	}
	defer res.Body.Close()
	var userByUserName userByUserNameRes
	// reading from response body until an error or EOF
	err = json.NewDecoder(res.Body).Decode(&userByUserName)
	if err != nil {
		log.GenericError(ctx.TContext, err, nil)
		return nil, err
	}

	return &userByUserName.Data, err
}

func userByID(ctx *customCtx, userID string) (*user, error) {
	endpoint := consulhelper.GetServiceNodes(ctx.TContext, l.ServiceID) + l.UserByID + userID

	res, err := hwrap.HTTPRequest(ctx.TContext, hwrap.RequestParams{
		Method:      http.MethodGet,
		URL:         endpoint,
		ContentType: com.AppJSON,
		Body:        nil,
	})
	if err != nil {
		err = fmt.Errorf("call to %s failed, error: %v", endpoint, err)
		log.GenericError(ctx.TContext, err, nil)
		return nil, err
	}

	defer res.Body.Close()
	var userByID userByIDRes
	// reading from response body until an error or EOF
	err = json.NewDecoder(res.Body).Decode(&userByID)
	if err != nil {
		log.GenericError(ctx.TContext, err, nil)
		return nil, err
	}

	return &userByID.Data, err
}
