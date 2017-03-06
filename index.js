import API from './api';
import {TEKION_DEALER_CREATE_REQUEST,TEKION_DEALER_CREATE_SUCCESS,TEKION_DEALER_CREATE_FAILURE} from './constants'

export function GetDealerByID(cfg,cfg.dealerId) {
   return {
     types: [TEKION_DEALER_CREATE_REQUEST,TEKION_DEALER_CREATE_SUCCESS,TEKION_DEALER_CREATE_FAILURE],
     promise: () => { return API.CreateCustomer(cfg,cfg.dealerId)}
   };
}

	