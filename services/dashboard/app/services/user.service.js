'use strict';

UserService.inject = ['$http'];

function UserService($http, UserServiceEndpoint,
  ApplicationServiceEndpoint, AuthService) {

  var _this = this;

  /*
  _this.getApplications = function() {
    return $http.get(ApplicationAPI + '/applications');
  };
  */

  // add authentication methods here

}

module.exports = UserService;
