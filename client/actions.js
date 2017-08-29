import { createAction } from 'redux-actions';
import {
  TEKION_DEALER_INFO_REQUEST,
  TEKION_DEALER_INFO_SUCCESS,
  TEKION_DEALER_INFO_FAILURE,
  TEKION_GET_FIXED_OPERATION_REQUEST,
  TEKION_GET_FIXED_OPERATION_SUCCESS,
  TEKION_GET_FIXED_OPERATION_FAILURE,
} from './constants';
import Services from './services';

const getDealerInfoRequest = createAction(TEKION_DEALER_INFO_REQUEST);
const getDealerInfoSuccess = createAction(TEKION_DEALER_INFO_SUCCESS);
const getDealerInfoFailure = createAction(TEKION_DEALER_INFO_FAILURE);

const getFixedOperationsRequest = createAction(
  TEKION_GET_FIXED_OPERATION_REQUEST,
);
const getFixedOperationsSuccess = createAction(
  TEKION_GET_FIXED_OPERATION_SUCCESS,
);
const getFixedOperationsFailure = createAction(
  TEKION_GET_FIXED_OPERATION_FAILURE,
);

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

export function getFixedOperationForDealer(config) {
  return async (dispatch) => {
    dispatch(getFixedOperationsRequest);
    const { error, response } = await Services.getFixedOperation(config);
    if (response) {
      dispatch(getFixedOperationsSuccess(response));
    }
    if (error) {
      dispatch(getFixedOperationsFailure(error));
    }
  };
}
