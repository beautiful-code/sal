'use strict';

function AuthInterceptor(UserServiceEndpoint, ApplicationServiceEndpoint, AuthService) {
  return {
    // automatically attach Authorization header
    request: function(config) {
      var token = AuthService.getToken();

      if (token) {
        if (config.url.indexOf(UserServiceEndpoint) === 0 || config.url.indexOf(ApplicationServiceEndpoint) === 0) {
          config.headers.Authorization = token;
        }
      }

      return config;
    },

    // If a token was sent back, save it
    response: function(res) {
      if(res.config.url.indexOf(UserServiceEndpoint) === 0 && res.data.token) {
        AuthService.saveToken(res.data.token);
      }

      return res;
    },
  };
}

module.exports = AuthInterceptor;
