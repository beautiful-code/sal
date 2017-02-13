'use strict';

UserService.$inject = [
  '$http',
  'UserServiceEndpoint'
];

function UserService($http, UserServiceEndpoint) {
  var _this = this;

  _this.register = function(user) {
    return $http.post(UserServiceEndpoint + '/register', {
      data: user
    });
  };

  _this.login = function(user) {
    return $http.post(UserServiceEndpoint + '/login', {
      data: user
    });
  };

  _this.getUser = function() {
    return $http.get(UserServiceEndpoint + '/user');
  };
}

module.exports = UserService;
