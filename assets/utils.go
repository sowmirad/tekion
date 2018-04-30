package assets

import (
	"fmt"

	"strings"

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
