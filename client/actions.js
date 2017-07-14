import { createAction } from 'redux-actions';
import {
  TEKION_DEALER_INFO_REQUEST,
  TEKION_DEALER_INFO_SUCCESS,
  TEKION_DEALER_INFO_FAILURE,
} from './constants';
import { dealerInfo } from './mock_dealer_data';

const getDealerInfoRequest = createAction(TEKION_DEALER_INFO_REQUEST);
const getDealerInfoSuccess = createAction(TEKION_DEALER_INFO_SUCCESS);
const getDealerInfoFailure = createAction(TEKION_DEALER_INFO_FAILURE);

export function getDealerInfo() {
  return async (dispatch, getState) => {
    dispatch(getDealerInfoRequest);
    dispatch(getDealerInfoSuccess({ dealerInfo }));
  };
}
