'use strict';

angular.module('dashboard')
  //.constant('UserServiceEndpoint', 'http://192.168.99.100:31767')
  .constant('UserServiceEndpoint', 'http://localhost:8080')
  .constant('ApplicationServiceEndpoint', 'http://localhost:8090')
 //.constant('ApplicationServiceEndpoint', 'http://192.168.99.100:30347')
  .service('AuthService', require('./auth.service'))
  .service('UserService', require('./user.service'))
  .service('ApplicationService', require('./application.service'));
