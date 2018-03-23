import ip from 'icepick';
import idx from 'idx';

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
  TEKION_GET_LABORTYPES_FOR_MAKE,
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
  cpDefautLaborType: '',
  ipDefautLaborType: '',
  wpDefautLaborType: '',
  cp_defaultLaborType: null,
  ip_defaultLaborType: null,
  wp_defaultLaborType: null,
});

function convertPaytypeForUI(paytypes) {
  const result = [];
  const paytypeList = paytypes || [];
  for (let index = 0; index < paytypeList.length; index++) {
    const payObj = {
      laborTypeID: '',
      code: '',
      description: '',
      key: '',
    };

    const itemObj = paytypeList[index];

    if (itemObj.description && itemObj.description.length > 0) {
      payObj.key = `${itemObj.code} | ${itemObj.description}`;
    } else {
      payObj.key = `${itemObj.code}`;
    }

    payObj.code = itemObj.code;
    payObj.description = itemObj.description;
    payObj.laborTypeID = itemObj.laborTypeID;

    result[index] = payObj;
  }

  return result;
}

function convertDefaultPayTypeForUI(payTypeObject) {
  const result = [];

  const payObj = {
    laborTypeID: '',
    code: '',
    description: '',
    key: '',
  };

  const itemObj = payTypeObject;
  if (itemObj.description && itemObj.description.length > 0) {
    payObj.key = `${itemObj.code} | ${itemObj.description}`;
  } else {
    payObj.key = `${itemObj.code}`;
  }
  payObj.code = itemObj.code;
  payObj.description = itemObj.description;
  payObj.laborTypeID = itemObj.laborTypeID;

  result[0] = payObj;

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
      const paytype = idx(action.payload, _ => _.data.payTypes);

      state = ip.set(state, 'fixedOperationStatus', 'success');
      state = ip.set(state, 'fixedOperationData', action.payload.data);
      return state;
    
    case TEKION_GET_LABORTYPES_FOR_MAKE: {
      let make = action.payload && action.payload.toLowerCase();
      let paytype = idx(state.fixedOperationData, _ => _.makePayTypes.makePayTypesMap[make]);
      if(!paytype) {
        make = idx(state.fixedOperationData, _ => _.makePayTypes.defaultMake);
        paytype = idx(state.fixedOperationData, _ => _.makePayTypes.makePayTypesMap[make]);
      }
      state = ip.set(
        state,
        'cpPayTypeList',
        convertPaytypeForUI(idx(paytype, _ => _.CustomerPay.laborTypes)),
      );
      state = ip.set(
        state,
        'wpPayTypeList',
        convertPaytypeForUI(idx(paytype, _ => _.WarrantyPay.laborTypes)),
      );
      state = ip.set(
        state,
        'ipPayTypeList',
        convertPaytypeForUI(idx(paytype, _ => _.InternalPay.laborTypes)),
      );

      state = ip.set(
        state,
        'cpDefautLaborType',
        convertDefaultPayTypeForUI(
          idx(paytype, _ => _.CustomerPay.defaultLaborType),
        ),
      );
      state = ip.set(
        state,
        'ipDefautLaborType',
        convertDefaultPayTypeForUI(
          idx(paytype, _ => _.InternalPay.defaultLaborType),
        ),
      );
      state = ip.set(
        state,
        'wpDefautLaborType',
        convertDefaultPayTypeForUI(
          idx(paytype, _ => _.WarrantyPay.defaultLaborType),
        ),
      );

      state = ip.set(
        state,
        'cp_defaultLaborType',
        idx(paytype, _ => _.CustomerPay.defaultLaborType),
      );
      state = ip.set(
        state,
        'ip_defaultLaborType',
        idx(paytype, _ => _.InternalPay.defaultLaborType),
      );
      state = ip.set(
        state,
        'wp_defaultLaborType',
        idx(paytype, _ => _.WarrantyPay.defaultLaborType),
      );
      return state;
    }

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
