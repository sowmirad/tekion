package dealerService

import (
	log "bitbucket.org/tekion/tbaas/log/v1"
	redis "bitbucket.org/tekion/tbaas/redisHelper/v2"
	enumCom "bitbucket.org/tekion/tenums/common"
	"fmt"
)

/********* redis cache related functions ************************/
func setCacheData(ctx *customCtx, payload interface{}, serviceID, collection string, identifier string) {
	err := redis.AddOrUpdateCache(ctx.Tenant, ctx.DealerID, serviceID, collection, identifier, payload, enumCom.RedisExpireTime)
	if err != nil {
		err = fmt.Errorf("failed to set %s cache data in collection % s for %s , error: %v", serviceID, collection, identifier, err)
		log.GenericError(ctx.TContext, err, nil)
	}
}

/********* redis cache related functons ends ************************/
