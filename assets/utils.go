package assets

import (
	"fmt"
	"strings"

	"bitbucket.org/tekion/tbaas/apiContext"
	log "bitbucket.org/tekion/tbaas/log/v1"
	mMgr "bitbucket.org/tekion/tbaas/mongoManager"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

func (arb *assetsReqBody) findQ() (bson.M, error) {
	findQ := bson.M{}
	if len(arb.Origins) > 0 {
		sliceToLower(arb.Origins)
		findQ["origins"] = bson.M{"$in": arb.Origins}
	}
	if len(arb.Makes) > 0 {
		sliceToLower(arb.Makes)
		findQ["makes"] = bson.M{"$in": arb.Makes}
	}
	if len(arb.DealerIDs) > 0 {
		findQ["dealerIDs"] = bson.M{"$in": arb.DealerIDs}
	}

	if len(findQ) == 0 {
		return nil, fmt.Errorf(" missing search params ")
	}

	return findQ, nil
}

func sliceToLower(s []string) {
	for i := range s {
		s[i] = strings.ToLower(s[i])
	}
}

func (arb *assetsReqBody) searchQ() bson.M {
	if len(arb.Modules) == 0 {
		return nil
	}

	searchQ := bson.M{}
	for _, mods := range arb.Modules {
		searchQ[mods] = 1
	}
	return searchQ
}

func (arb *assetsReqBody) findAsset(ctx apiContext.TContext, findQ, searchQ bson.M) ([]assets, error) {
	var err error
	assets := make([]assets, 0, 0)
	if err = mMgr.ReadAll(ctx.Tenant, assetCol, findQ, searchQ, &assets); err != nil {
		err = errors.Wrap(err, " failed to read assets from db ")
		return nil, err
	}
	if len(assets) == 0 {
		msg := fmt.Sprintf(" no assets found in db, findQ:%+v, request:%+v, looking for default asset ", findQ, arb)
		log.GenericInfo(ctx, msg, nil)

		findQ = arb.findQDefault()
		if err = mMgr.ReadAll(ctx.Tenant, assetCol, findQ, searchQ, &assets); err != nil {
			err = errors.Wrap(err, " failed to read default assets from db ")
			return nil, err
		}
	}
	return assets, err
}

func (arb *assetsReqBody) findQDefault() bson.M {
	findQ := bson.M{}

	findQ["dealerIDs"] = bson.M{"$in": arb.DealerIDs}
	findQ["origins"] = bson.M{"$eq": []string{}}
	findQ["makes"] = bson.M{"$eq": []string{}}

	return findQ
}

func (arb *assetsReqBody) validateRes(assets []assets, findQ bson.M) error {
	if len(assets) == 0 {
		return fmt.Errorf(" no asset found in db, findQ:%+v, request:%+v ", findQ, arb)
	}

	if len(assets) > 1 {
		return fmt.Errorf(" multiple assets returned from db, findQ:%+v, request:%+v ", findQ, arb)
	}
	return nil
}
