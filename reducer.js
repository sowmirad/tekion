'use strict';

import {TEKION_DEALER_CREATE_REQUEST,TEKION_DEALER_CREATE_SUCCESS,TEKION_DEALER_CREATE_FAILURE} from './constants'

var getDealerByIdState = {
  "result": {
    "meta": {
      "code": 200,
      "msg": ""
    },
    "data": {
    }
  },
  "error": "not found"
};

export function GetDealerByIDAPIReducer(state = getDealerByIdState, action ) {
  switch (action.type) {
      case TEKION_DEALER_CREATE_REQUEST:
      {
         return {
            ...state,
            'error': null,
            'result': null
        };
      }
      break;
      case TEKION_DEALER_CREATE_SUCCESS:
      {
         return {
            ...state,
            'error': null,
            'result': action.result
          };
      }
      break;
      case TEKION_DEALER_CREATE_FAILURE:
      {
          return {
            ...state,
            'error': action.error,
            'result': null
          };
      }
      break;
      default:
        return state;
  }
}
}



