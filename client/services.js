import { GET, POST } from './api';

export default {
  getDealerInfo: GET('/tdealer/dealer'),
  getFixedOperation: GET('/tdealer/fixedoperation'),
  getDealerList: POST('/tdealer/dealers'),
  updateDealerInfo: POST('/tdealer/dealer'),
  createDealer: POST('tdealer/dealer'),
};
