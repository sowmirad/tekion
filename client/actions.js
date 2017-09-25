import { createAction } from 'redux-actions';
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

const getDealerListRequest = createAction(TEKION_DEALER_LIST_REQUEST);
const getDealerListSuccess = createAction(TEKION_DEALER_LIST_SUCCESS);
const getDealerListFailure = createAction(TEKION_DEALER_LIST_FAILURE);

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

export function setDealerInfo(dealerInfo) {
  return async (dispatch) => {
      dispatch(getDealerInfoSuccess({ dealerInfo }));
    }
}

export function getDealerList(config) {
  return async (dispatch) => {
    dispatch(getDealerListRequest);
    const { error, response } = await Services.getDealerList(config);
    if (response) {
      dispatch(getDealerListSuccess({ dealerInfo: response.data }));
    }
    if (error) {
      dispatch(getDealerListFailure(error));
    }
  };
}
