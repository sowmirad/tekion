import ip from 'icepick';
import {
  TEKION_DEALER_INFO_REQUEST,
  TEKION_DEALER_INFO_SUCCESS,
  TEKION_DEALER_INFO_FAILURE,
  TEKION_GET_FIXED_OPERATION_REQUEST,
  TEKION_GET_FIXED_OPERATION_SUCCESS,
  TEKION_GET_FIXED_OPERATION_FAILURE,
} from './constants';

const initialState = ip.freeze({
  dealerInfo: null,
  requestStatus: null,
  fixedOperationStatus: null,
  fixedOperationData: null,
});

export default function (state = initialState, action) {
  switch (action.type) {
    case TEKION_DEALER_INFO_REQUEST:
      state = ip.set(state, 'dealerInfo', null);
      state = ip.set(state, 'requestStatus', 'fetching');
      return state;

    case TEKION_DEALER_INFO_SUCCESS:
      state = ip.set(state, 'dealerInfo', action.payload.dealerInfo);
      state = ip.set(state, 'requestStatus', 'success');
      return state;

    case TEKION_DEALER_INFO_FAILURE:
      state = ip.set(state, 'dealerInfo', null);
      state = ip.set(state, 'requestStatus', 'failed');
      return state;

    case TEKION_GET_FIXED_OPERATION_REQUEST:
      state = ip.set(state, 'fixedOperationStatus', 'fetching');
      state = ip.set(state, 'fixedOperationData', null);
      return state;

    case TEKION_GET_FIXED_OPERATION_SUCCESS:
      state = ip.set(state, 'fixedOperationStatus', 'success');
      state = ip.set(state, 'fixedOperationData', action.payload.data);
      return state;

    case TEKION_GET_FIXED_OPERATION_FAILURE:
      state = ip.set(state, 'fixedOperationData', 'failed');
      state = ip.set(state, 'fixedOperationData', null);
      return state;

    default:
      return state;
  }
}
