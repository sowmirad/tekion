'use strict';

export default class API {
  static setHeaders(cfg) {
  console.log('header value is '+JSON.stringify(cfg));
    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    // Add tenantid in header
    if(cfg.tenantId){
        myHeaders.append("tenantName", cfg.tenantId);
    }
    // Add dealerid in header
    if(cfg.dealerId){
        myHeaders.append("dealerId", cfg.dealerId);
    }
    // Add access_token in header
    if(cfg.access_token){
        myHeaders.append("tekion-api-token", cfg.access_token);
    }
    return myHeaders;
  }

  static GetDealerByID(cfg, cfg.dealerId) {
       return new Promise((resolve, reject) => {
             var myInit = { method: 'GET',
                headers: API.setHeaders(cfg),
                mode: 'cors',
                cache: 'default',
             };
             cfg.isDebug && console.log(cfg.baseURL+'/tdealer/getDealerById');
             return fetch(cfg.baseURL+'/tdealer/getDealerById' , myInit).then(function(response) {
                 // fetch call came back with some response
                 cfg.isDebug && console.log("successful fetch")
                 response.json().then(function(data){
                 cfg.isDebug && console.log(data)

                 // check when the username is an email
                 if (response.ok && data.meta.code === 200)
                   resolve(data);
                 // check when the user name is not email and there
                 // is group name
                 else if(response.ok && data.meta.code === 100)
                   resolve(data.meta);
                 else
                   reject(data.meta.msg);
               },
               function(error) {
               cfg.isDebug && console.log('failed json decode')
                 cfg.isDebug && console.log(error)
                   // failed to change response to json
                   reject(error);
               });
             }, function(error) {
               // fatal error
               cfg.isDebug && console.log('failed fetch')
               cfg.isDebug && console.log(error)
               reject(error);
             });
       });
   }
}