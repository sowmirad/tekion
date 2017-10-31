import * as actions from './actions';
import dealerReducer from './dealerReducer';

module.exports = {
  dealerReducer,
  ...actions,
};
