import ip from 'icepick';
import {
  TEKION_DEALER_INFO_REQUEST,
  TEKION_DEALER_INFO_SUCCESS,
  TEKION_DEALER_INFO_FAILURE,
} from './constants';

const initialState = ip.freeze({
  dealerInfo: null,
  requestStatus: null,
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

    default:
      return state;
  }
}
