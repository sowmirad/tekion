import ip from 'icepick';
import {
  TEKION_DEALER_INFO_REQUEST,
  TEKION_DEALER_INFO_SUCCESS,
  TEKION_DEALER_INFO_FAILURE,
  TEKION_GET_FIXED_OPERATION_REQUEST,
  TEKION_GET_FIXED_OPERATION_SUCCESS,
  TEKION_GET_FIXED_OPERATION_FAILURE,
  TEKION_DEALER_LIST_REQUEST,
  TEKION_DEALER_LIST_SUCCESS,
  TEKION_DEALER_LIST_FAILURE,
  TEKION_UPDATE_DEALER_INFO_REQUEST,
  TEKION_UPDATE_DEALER_INFO_SUCCESS,
  TEKION_UPDATE_DEALER_INFO_FAILURE,
  TEKION_CREATE_DEALER_REQUEST,
  TEKION_CREATE_DEALER_SUCCESS,
  TEKION_CREATE_DEALER_FAILURE,
  TEKION_DEALER_RESET_STATE,
} from './constants';

const initialState = ip.freeze({
  dealerInfo: null,
  requestStatus: null,
  fixedOperationStatus: null,
  fixedOperationData: null,
  dealerList: null,
  dealerListStatus: null,
  dealerData: null,
  updateDealerDataStatus: null,
  newDealerData: null,
  newDealerDataStatus: null,
  cpPayTypeList: [],
  ipPayTypeList: [],
  wpPayTypeList: [],
});

function convertPaytypeForUI(paytypeList) {
  console.log('inside', paytypeList);

  const result = [];
  const payObj = {
    laborTypeID: '',
    code: '',
    description: '',
    key: '',
  };

  for (let index = 0; index < paytypeList.length; index++) {
    const itemObj = paytypeList[index];

    payObj.key = `${itemObj.code} |${itemObj.description}`;
    payObj.code = itemObj.code;
    payObj.description = itemObj.description;
    payObj.laborTypeID = itemObj.laborTypeID;

    result.push(payObj);
  }

  return result;
}

export default function (state = initialState, action) {
  switch (action.type) {
    case TEKION_DEALER_INFO_REQUEST:
      // state = ip.set(state, 'dealerInfo', null);
      state = ip.set(state, 'requestStatus', 'fetching');
      return state;

    case TEKION_DEALER_INFO_SUCCESS:
      state = ip.set(state, 'dealerInfo', action.payload.dealerInfo);
      state = ip.set(state, 'requestStatus', 'success');
      return state;

    case TEKION_DEALER_INFO_FAILURE:
      // state = ip.set(state, 'dealerInfo', null);
      state = ip.set(state, 'requestStatus', 'failed');
      return state;

    case TEKION_GET_FIXED_OPERATION_REQUEST:
      state = ip.set(state, 'fixedOperationStatus', 'fetching');
      // state = ip.set(state, 'fixedOperationData', null);
      return state;

    case TEKION_GET_FIXED_OPERATION_SUCCESS:
      const paytype = action.payload.data.payTypes;

      state = ip.set(state, 'fixedOperationStatus', 'success');
      state = ip.set(state, 'fixedOperationData', action.payload.data);
      state = ip.set(
        state,
        'cpPayTypeList',
        convertPaytypeForUI(paytype.CustomerPay.laborTypes),
      );
      state = ip.set(
        state,
        'wpPayTypeList',
        convertPaytypeForUI(paytype.WarrantyPay.laborTypes),
      );
      state = ip.set(
        state,
        'ipPayTypeList',
        convertPaytypeForUI(paytype.InternalPay.laborTypes),
      );
      return state;

    case TEKION_GET_FIXED_OPERATION_FAILURE:
      state = ip.set(state, 'fixedOperationStatus', 'failed');
      // state = ip.set(state, 'fixedOperationData', null);
      return state;

    case TEKION_DEALER_LIST_REQUEST:
      state = ip.set(state, 'dealerList', null);
      state = ip.set(state, 'dealerListStatus', 'fetching');
      return state;

    case TEKION_DEALER_LIST_SUCCESS:
      state = ip.set(state, 'dealerList', action.payload.dealerList);
      state = ip.set(state, 'dealerListStatus', 'success');
      return state;

    case TEKION_DEALER_LIST_FAILURE:
      state = ip.set(state, 'dealerList', null);
      state = ip.set(state, 'dealerListStatus', 'failed');
      return state;

    case TEKION_UPDATE_DEALER_INFO_REQUEST:
      state = ip.set(state, 'dealerData', null);
      state = ip.set(state, 'updateDealerDataStatus', 'fetching');
      return state;

    case TEKION_UPDATE_DEALER_INFO_SUCCESS:
      state = ip.set(state, 'dealerData', action.payload.dealerInfo);
      state = ip.set(state, 'updateDealerDataStatus', 'success');
      return state;

    case TEKION_UPDATE_DEALER_INFO_FAILURE:
      state = ip.set(state, 'dealerData', null);
      state = ip.set(state, 'updateDealerDataStatus', 'failed');
      return state;

    case TEKION_CREATE_DEALER_REQUEST:
      state = ip.set(state, 'newDealerData', null);
      state = ip.set(state, 'newDealerDataStatus', 'fetching');
      return state;

    case TEKION_CREATE_DEALER_SUCCESS:
      state = ip.set(state, 'newDealerData', action.payload.dealerInfo);
      state = ip.set(state, 'newDealerDataStatus', 'success');
      return state;

    case TEKION_CREATE_DEALER_FAILURE:
      state = ip.set(state, 'newDealerData', null);
      state = ip.set(state, 'newDealerDataStatus', 'failed');
      return state;

    case TEKION_DEALER_RESET_STATE:
      let targetState = JSON.parse(JSON.stringify(action.payload));
      targetState = ip.freeze(targetState);
      return targetState;

    default:
      return state;
  }
}
