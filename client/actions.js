import { createAction } from 'redux-actions';
import {
  TEKION_DEALER_INFO_REQUEST,
  TEKION_DEALER_INFO_SUCCESS,
  TEKION_DEALER_INFO_FAILURE,
} from './constants';
import Services from './services';

const getDealerInfoRequest = createAction(TEKION_DEALER_INFO_REQUEST);
const getDealerInfoSuccess = createAction(TEKION_DEALER_INFO_SUCCESS);
const getDealerInfoFailure = createAction(TEKION_DEALER_INFO_FAILURE);

export function getDealerInfo(config) {
  return async (dispatch) => {
    dispatch(getDealerInfoRequest);
    const { error, response } = await Services.getDealerInfo(config);
    if (response) {
      dispatch(getDealerInfoSuccess({ dealerInfo: response.data }));
    }
    if (error) {
      dispatch(getDealerInfoFailure(error));
    }
  };
}
