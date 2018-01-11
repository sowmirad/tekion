import querystring from 'querystring';
import axios from 'axios';

const makeMethod = (method, hasBody) => urlTemplate => (config, data) => {
  let url = config.baseURL + urlTemplate;
  for (const tag of url.match(/:\w+/g) || []) {
    let value = data[tag.slice(1)];
    if (value === undefined) {
      console.warn('Warning: calling', method, 'without', tag);
      value = '';
    }
    url = url.replace(tag, encodeURIComponent(data[tag.slice(1)]));
    delete data[tag.slice(1)];
  }

  const headers = {};
  if (config) {
    headers.tenantName = config.tenantName;
    headers.dealerId = config.dealerId;
    headers['tekion-api-token'] = config.access_token;
    if (config.correlationId) {
      headers.correlationId = config.correlationId;
    }
  }

  if (config && config.clientId) {
    headers.clientID = config.clientId;
  }

  if (!hasBody) {
    const qs = querystring.stringify(data);
    if (qs) {
      url += (url.indexOf('?') >= 0 ? '&' : '?') + qs;
    }
  }

  let axiosConfig = { method, headers, url };

  if (hasBody) {
    axiosConfig = { ...axiosConfig, data };
  }

  if (config.isDebug) {
    console.log(axiosConfig);
  }

  return axios(axiosConfig)
    .then((response) => {
      if (config.isDebug) {
        console.log(response);
      }
      return { error: null, response: response.data };
    })
    .catch((error) => {
      if (config.isDebug) {
        console.log(error);
      }
      if (error.response) {
        return { error: error.response, response: null };
      }
      return { error: { code: 9999, msg: 'User is Offline' }, response: null };
    });
};

const GET = makeMethod('GET');
const DELETE = makeMethod('DELETE');
const POST = makeMethod('POST', true);
const PUT = makeMethod('PUT', true);

export { GET, DELETE, POST, PUT };
